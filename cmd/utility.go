package main

import (
	"CSVtoSQL/internal/db"
	"CSVtoSQL/internal/env"
	"CSVtoSQL/internal/importer"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func PromptDBType(reader *bufio.Reader) string {
	fmt.Print("Выберите тип базы данных (mysql): ")
	dbType, _ := reader.ReadString('\n')
	dbType = strings.TrimSpace(strings.ToLower(dbType))
	return dbType
}

func PromptMySQLParams(reader *bufio.Reader) (host, port, user, password, dbname, table, csvFile, ow string) {
	fmt.Print("Введите хост:")
	host, _ = reader.ReadString('\n')
	fmt.Print("Введите порт:")
	port, _ = reader.ReadString('\n')
	fmt.Print("Введите имя пользователя: ")
	user, _ = reader.ReadString('\n')
	fmt.Print("Введите пароль: ")
	password, _ = reader.ReadString('\n')
	fmt.Print("Введите имя базы данных: ")
	dbname, _ = reader.ReadString('\n')
	fmt.Print("Введите имя таблицы: ")
	table, _ = reader.ReadString('\n')
	fmt.Print("Введите имя CSV файла (файл должен находиться рядом с приложением): ")
	csvFile, _ = reader.ReadString('\n')
	fmt.Print("Перезаписывать данные в таблице при конфликте? (y/n): ")
	overwrite, _ := reader.ReadString('\n')
	ow = "false"
	if strings.TrimSpace(strings.ToLower(overwrite)) == "y" {
		ow = "true"
	}
	return
}

func SaveEnv(dbType, host, port, user, password, dbname, table, csvFile, ow string) {
	vars := map[string]string{
		"DB_TYPE":               dbType,
		"HOST":                  strings.TrimSpace(host),
		"PORT":                  strings.TrimSpace(port),
		"USER":                  strings.TrimSpace(user),
		"PASSWORD":              strings.TrimSpace(password),
		"DB":                    strings.TrimSpace(dbname),
		"TABLE":                 strings.TrimSpace(table),
		"CSV":                   strings.TrimSpace(csvFile),
		"OVERWRITE_ON_CONFLICT": ow,
	}
	err := env.WriteEnv(vars, ".env")
	if err != nil {
		fmt.Println("Ошибка записи .env:", err)
		os.Exit(1)
	}
	fmt.Println("Параметры сохранены в .env")
}

func LoadEnvOrExit() {
	err := env.LoadEnv(".env")
	if err != nil {
		fmt.Println("Ошибка загрузки .env:", err)
		os.Exit(1)
	}
}

func Run() {
	dbType := os.Getenv("DB_TYPE")

	var dbService *db.DBService
	var err error

	switch dbType {
	case "mysql":
		dbService, err = db.NewMySQLService()
		if err != nil {
			fmt.Println("Ошибка подключения", err)
			os.Exit(1)
		}
	default:
		fmt.Println("Данная DB пока не поддерживается")
		os.Exit(1)
	}

	defer dbService.Close()

	if err = dbService.Ping(); err != nil {
		fmt.Println("Не удалось подключиться", err)
		os.Exit(1)
	}

	fmt.Println("Успешное подключение!")

	csvPath := os.Getenv("CSV")
	importer, err := importer.NewImporter(csvPath, dbService)
	if err != nil {
		fmt.Println("Ошибка открытия CSV:", err)
		os.Exit(1)
	}

	table := os.Getenv("TABLE")
	overwrite := os.Getenv("OVERWRITE_ON_CONFLICT") == "true"
	err = importer.ImportAll(table, overwrite)
	if err != nil {
		fmt.Println("Ошибка импорта:", err)
		os.Exit(1)
	}
	fmt.Println("Импорт завершён успешно!")
}
