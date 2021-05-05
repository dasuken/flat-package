package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func NewDBConn() (*sql.DB, error) {
	return sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/product")
}