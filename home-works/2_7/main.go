package main

import (
	"2_7/argon2"
	"2_7/db"
	"2_7/input"
	"fmt"

	_ "modernc.org/sqlite"
)

func auth(pare input.Pare) {
	user, err := db.FindUserByLogin(pare.Login)
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

func main() {
	db.AddDefaultUsers()

	err, pare := input.InputAuth()
	if err != nil {
		fmt.Println(err)
		return
	}

	auth(pare)
}
