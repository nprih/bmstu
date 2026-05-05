package main

import "fmt"

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
		return fmt.Errorf("Error while calling f3: %w", err)
	}
	return nil
}
func f3() error {
	return fmt.Errorf("I'm error in f3")
}
func main() {
	err := f1()
	if err != nil {
		fmt.Println(err)
	}
}
