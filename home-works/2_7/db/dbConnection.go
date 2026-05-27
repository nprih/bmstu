package db

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

func DbConnection() (error, *sql.DB) {
	db, err := sql.Open("sqlite", "users")
	if err != nil {
		return err, db
	}
	return nil, db
}
