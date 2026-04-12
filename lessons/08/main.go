package main

import "fmt"

func main() {
	defer fmt.Println("This is defer #1")
	defer fmt.Println("This is defer #2")
	defer fmt.Println("This is defer #3")
	panic("dssddsdsds")
	fmt.Println("main function")
}
