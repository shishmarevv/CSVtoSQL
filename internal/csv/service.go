package csv

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"
)

func NewCSVReader(filename string) (*CSVReader, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("ошибка открытия файла: %w", err)
	}
	scanner := bufio.NewScanner(f)
	rowCount := -1 // -1 чтобы не считать заголовок
	for scanner.Scan() {
		rowCount++
	}
	f.Close()
	if err = scanner.Err(); err != nil {
		return nil, fmt.Errorf("ошибка сканирования файла: %w", err)
	}

	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("ошибка открытия файла: %w", err)
	}
	reader := csv.NewReader(file)
	head, err := reader.Read()
	head = removeBOM(head)
	if err != nil {
		file.Close()
		return nil, fmt.Errorf("ошибка чтения заголовка: %w", err)
	}
	return &CSVReader{file: file, reader: reader, RowCount: rowCount, Header: head}, nil
}

func (c *CSVReader) ReadRow() ([]string, error) {
	row, err := c.reader.Read()
	if err == io.EOF {
		return nil, io.EOF
	}
	if err != nil {
		return nil, fmt.Errorf("ошибка чтения строки: %w", err)
	}
	return removeBOM(row), nil
}

func (c *CSVReader) Close() error {
	return c.file.Close()
}

func removeBOM(fields []string) []string {
	if len(fields) == 0 {
		return fields
	}
	fields[0] = strings.TrimPrefix(fields[0], "\uFEFF")
	return fields
}
