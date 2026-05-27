package main

import (
	"2_7/auth"
	"2_7/db"
	"2_7/input"
	"fmt"

	_ "modernc.org/sqlite"
)

func main() {
	db.AddDefaultUsers()

	err, pare := input.InputAuth()
	if err != nil {
		fmt.Println(err)
		return
	}

	auth.Auth(pare)
}
