package main

import (
	"fmt"
	"os"
)

type User struct {
	Email    string
	Password string
	Age      int
}

func main() {
	file, err := os.OpenFile("test.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	err = os.MkdirAll("Work/test/test2/test3", 0777)
	if err != nil {
		fmt.Println(err)
	}
}
