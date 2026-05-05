package main

import (
	"fmt"
)

type newError struct{}

func (newerr newError) Error() string {
	return "This is new Error"
}

func someFunc() error {
	var NextGenError newError
	return NextGenError
}

func main() {
	err := fmt.Errorf("Some %s", "Error") // error - interface -> Error()
	//err = errors.New(fmt.Sprintf("Another %s", "Error"))
	fmt.Println(someFunc())

}
