package dbConnection

import "database/sql"

func DbConnection() (error, *sql.DB) {
	db, err := sql.Open("sqlite", "users")
	if err != nil {
		return err, db
	}
	return nil, db
}
