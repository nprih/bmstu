package auth

import (
	"2_7/argon2"
	"2_7/db"
	"2_7/input"
	"fmt"
)

func Auth(pare input.Pare) {
	user, err := db.FindUserByLogin(pare.Login)
	if err != nil {
		fmt.Println(err)
		return
	}
	ok, _ := argon2.CheckPasswordArgon2(pare.Password, user.Password)
	if !ok {
		fmt.Println("\nПароль неверный")
		return
	}
	fmt.Println("\nПароль верный")
}
