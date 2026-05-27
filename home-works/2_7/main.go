package main

import (
	"2_7/auth"
	"2_7/db"
	"2_7/input"
	"fmt"
)

func main() {
	db.AddDefaultUsers()

	err, pair := input.InputAuth()
	if err != nil {
		fmt.Println(err)
		return
	}

	auth.Auth(pair)
}
