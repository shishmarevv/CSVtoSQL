package importer

import (
	"CSVtoSQL/internal/csv"
	"CSVtoSQL/internal/db"
	"io"
	"strings"
)

func NewImporter(csvPath string, db *db.DBService) (*Importer, error) {
	reader, err := csv.NewCSVReader(csvPath)
	if err != nil {
		return nil, err
	}
	return &Importer{
		CSV:    reader,
		DB:     db,
		Header: reader.Header,
	}, nil
}

func (imp *Importer) ImportAll(table string, overwrite bool) error {
	defer imp.CSV.Close()
	for {
		row, err := imp.CSV.ReadRow()
		if err != nil {
			if err.Error() == "EOF" || err == io.EOF {
				break
			}
			return err
		}
		err = imp.ImportRow(table, row, overwrite)
		if err != nil {
			return err
		}
	}
	return nil
}

func (imp *Importer) ImportRow(table string, row []string, overwrite bool) error {
	columns := imp.Header
	placeholders := make([]string, len(columns))
	for i := range placeholders {
		placeholders[i] = "?"
	}
	query := ""
	if overwrite {
		query = "INSERT INTO `" + table + "` (" + joinColumns(columns) + ") VALUES (" + joinPlaceholders(placeholders) + ") ON DUPLICATE KEY UPDATE " + buildUpdateSet(columns) + ";"
	} else {
		query = "INSERT IGNORE INTO `" + table + "` (" + joinColumns(columns) + ") VALUES (" + joinPlaceholders(placeholders) + ");"
	}
	_, err := imp.DB.DB.Exec(query, toInterfaceSlice(row)...)
	return err
}

func joinColumns(cols []string) string {
	return "`" + strings.Join(cols, "`, `") + "`"
}

func joinPlaceholders(ph []string) string {
	return strings.Join(ph, ", ")
}

func buildUpdateSet(cols []string) string {
	set := make([]string, len(cols))
	for i, col := range cols {
		set[i] = "`" + col + "`=VALUES(`" + col + "`)"
	}
	return strings.Join(set, ", ")
}

func toInterfaceSlice(strs []string) []interface{} {
	res := make([]interface{}, len(strs))
	for i, v := range strs {
		res[i] = v
	}
	return res
}
