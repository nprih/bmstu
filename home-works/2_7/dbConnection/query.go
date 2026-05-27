package dbConnection

import (
	"2_7/functions"
	"database/sql"
	"fmt"
	"log"
	"strings"
)

func ClearDb(db *sql.DB) {
	_, err := db.Exec("DELETE FROM users")
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec("DELETE FROM sqlite_sequence WHERE name = 'users'")
	if err != nil {
		log.Fatal(err)
	}
}

func AddDefaultUsers() {
	err, db := DbConnection()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	ClearDb(db)

	var placeholders []string
	var args []interface{}
	for _, user := range functions.UsersHashingPass() {
		placeholders = append(placeholders, "(?, ?)")
		args = append(args, user.Email, user.Password)
	}

	query := `INSERT INTO users (email, password) VALUES ` + strings.Join(placeholders, ", ")

	_, err = db.Exec(query, args...)
	if err != nil {
		log.Fatal(err)
	}
}
