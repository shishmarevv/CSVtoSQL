package main

import (
	"bufio"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	dbType := PromptDBType(reader)
	host, port, user, password, dbname, table, csvFile, ow := PromptMySQLParams(reader)
	SaveEnv(dbType, host, port, user, password, dbname, table, csvFile, ow)

	LoadEnvOrExit()

	Run()
}
