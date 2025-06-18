package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func NewMySQLService() (*DBService, error) {
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	user := os.Getenv("USER")
	password := os.Getenv("PASSWORD")
	dbname := os.Getenv("DB")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, dbname)
	db, err := sql.Open("db", dsn)
	if err != nil {
		return nil, err
	}
	return &DBService{DB: db}, nil
}

func (s *DBService) Ping() error {
	return s.DB.Ping()
}

func (s *DBService) Close() error {
	return s.DB.Close()
}
