package main

import (
	"CSVtoSQL/internal/env"
	"bufio"
	"fmt"
	"os"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	err := env.LoadEnv(".env")
	var skip bool
	if err != nil {
		fmt.Println("Error loading .env file")
		skip = false
	} else {
		skip = EnvIsFilled(reader)
	}
	if !skip {
		dbType := PromptDBType(reader)
		host, port, user, password, dbname, table, csvFile, ow := PromptMySQLParams(reader)
		SaveEnv(dbType, host, port, user, password, dbname, table, csvFile, ow)
	}

	LoadEnvOrExit()

	Run()
}
