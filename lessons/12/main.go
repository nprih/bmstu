package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type User struct {
	Email    string
	Password string
	Age      int
}

func main() {
	file, err := os.OpenFile("test.json", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	admin := User{Email: "admin@mail.ru", Password: "123456", Age: 18}
	data, err := json.Marshal(admin)
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = file.Write(data)
	if err != nil {
		fmt.Println(err)
		return
	}
}
