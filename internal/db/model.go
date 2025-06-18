package db

import (
	"database/sql"
)

type DBService struct {
	DB *sql.DB
}
