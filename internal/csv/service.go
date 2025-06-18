package csv

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

func NewCSVReader(filename string) (*CSVReader, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("ошибка открытия файла: %w", err)
	}
	reader := csv.NewReader(file)
	head, err := reader.Read()
	if err != nil {
		file.Close()
		return nil, fmt.Errorf("ошибка чтения заголовка: %w", err)
	}
	return &CSVReader{file: file, reader: reader, Header: head}, nil
}

func (c *CSVReader) ReadRow() ([]string, error) {
	row, err := c.reader.Read()
	if err == io.EOF {
		return nil, io.EOF
	}
	if err != nil {
		return nil, fmt.Errorf("ошибка чтения строки: %w", err)
	}
	return row, nil
}

func (c *CSVReader) Close() error {
	return c.file.Close()
}
