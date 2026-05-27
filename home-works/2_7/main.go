package main

import (
	"2_7/argon2"
	"2_7/dbConnection"
	"2_7/input"
	"2_7/models"
	"errors"
	"fmt"

	_ "modernc.org/sqlite"
)

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

func findUserByLogin(login string) (models.User, error) {
	err, db := dbConnection.DbConnection()
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

func main() {
	dbConnection.AddDefaultUsers()

	err, pare := input.InputAuth()
	if err != nil {
		fmt.Println(err)
		return
	}

	auth(pare)
}
