package infra

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func ConnectToDatabase() (db *sql.DB, err error) {
	db, err = sql.Open("sqlite3", "data.db")
	if err != nil {
		return
	}

	return
}
