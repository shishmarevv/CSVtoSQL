package importer

import (
	"CSVtoSQL/internal/csv"
	"CSVtoSQL/internal/db"
)

type Importer struct {
	CSV    *csv.CSVReader
	DB     *db.DBService
	Header []string
}
