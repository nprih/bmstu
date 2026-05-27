package main

import (
	"2_7/argon2"
	"2_7/dbConnection"
	"2_7/input"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strings"

	_ "modernc.org/sqlite"
)

type User struct {
	Id       int64
	Email    string
	Password string
}

func getDefaultUsers() []User {
	return []User{
		{
			Email:    "first@test.ru",
			Password: "123321",
		},
		{
			Email:    "second@test.ru",
			Password: "123qwe",
		},
		{
			Email:    "third@test.ru",
			Password: "ASD123",
		},
	}
}

func usersHashingPass() []User {
	users := getDefaultUsers()
	for i, user := range users {
		users[i].Password, _ = argon2.HashPasswordArgon2(user.Password)
	}
	return users
}

func addDefaultUsers() {
	err, db := dbConnection.DbConnection()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	clearDb(db)

	var placeholders []string
	var args []interface{}
	for _, user := range usersHashingPass() {
		placeholders = append(placeholders, "(?, ?)")
		args = append(args, user.Email, user.Password)
	}

	query := `INSERT INTO users (email, password) VALUES ` + strings.Join(placeholders, ", ")

	_, err = db.Exec(query, args...)
	if err != nil {
		log.Fatal(err)
	}
}

func clearDb(db *sql.DB) {
	_, err := db.Exec("DELETE FROM users")
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec("DELETE FROM sqlite_sequence WHERE name = 'users'")
	if err != nil {
		log.Fatal(err)
	}
}

func auth(pare input.Pare) {
	user, err := findUserByLogin(pare.Login)
	if err != nil {
		fmt.Println(err)
		return
	}

	ok, _ := argon2.CheckPasswordArgon2(pare.Password, user.Password)
	if !ok {
		fmt.Println("Пароль неверный")
		return
	}
	fmt.Println("Пароль верный")
}

func findUserByLogin(login string) (User, error) {
	err, db := dbConnection.DbConnection()
	if err != nil {
		fmt.Println(err)
		return User{}, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM users WHERE email = $1", login)
	if err != nil {
		fmt.Println(err)
		return User{}, err
	}
	defer rows.Close()

	users := []User{}
	for rows.Next() {
		user := User{}
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

	return User{}, errors.New(fmt.Sprintf("\nПользователь \"%s\" не найден", login))
}

func main() {
	addDefaultUsers()

	err, pare := input.InputAuth()
	if err != nil {
		fmt.Println(err)
		return
	}

	auth(pare)
}
