package csv

import (
	"encoding/csv"
	"os"
)

type CSVReader struct {
	file     *os.File
	reader   *csv.Reader
	RowCount int
	Header   []string
}
