package main

import (
	"errors"
	"fmt"
)

var errInF3 = fmt.Errorf("I'm error in f3")
var errInF2 = fmt.Errorf("Error while calling f3: %w", errInF3)

func f1() error {
	err := f2()
	if err != nil {
		return fmt.Errorf("Error while calling f2: %w", err)
	}
	return nil
}
func f2() error {
	err := f3() //
	if err != nil {
		return errInF2
	}
	return nil
}
func f3() error {
	return errInF3
}
func main() {
	err := f1()

	if errors.Is(err, errInF2) {
		fmt.Println("Yes, it is")
	}
}
