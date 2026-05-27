package db

import (
	"2_7/functions"
	"2_7/models"
	"database/sql"
	"errors"
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

func FindUserByLogin(login string) (models.User, error) {
	err, db := DbConnection()
	if err != nil {
		fmt.Println(err)
		return models.User{}, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM users WHERE email = $1", login)
	if err != nil {
		fmt.Println(err)
		return models.User{}, err
	}
	defer rows.Close()

	users := []models.User{}
	for rows.Next() {
		user := models.User{}
		err := rows.Scan(&user.Id, &user.Email, &user.Password)
		if err != nil {
			fmt.Println(err)
			continue
		}
		users = append(users, user)
	}
	if len(users) != 0 {
		return users[0], nil
	}

	return models.User{}, errors.New(fmt.Sprintf("\nПользователь \"%s\" не найден", login))
}
