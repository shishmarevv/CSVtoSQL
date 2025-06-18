package env

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

func WriteEnv(vars map[string]string, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	for k, v := range vars {
		line := fmt.Sprintf("%s=%s\n", k, v)
		_, err := file.WriteString(line)
		if err != nil {
			return err
		}
	}
	return nil
}

func LoadEnv(path string) error {
	return godotenv.Load(path)
}
