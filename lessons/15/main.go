package main

import (
	"errors"
	"fmt"
)

type MyError struct {
	Code int
	Text string
}

func (e *MyError) Error() string {
	return e.Text
}

func doSomething() error {
	return &MyError{Code: 404, Text: "Not found"}
}

func main() {
	err := doSomething()
	var customError *MyError
	if errors.As(err, &customError) {
		fmt.Println(customError.Code)
		fmt.Println(customError.Text)
	}
}
