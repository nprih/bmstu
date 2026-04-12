package main

import "fmt"

func WorkingWithDB() {
	//Open DB no logic
	//Some logic
	//Close DB connection
}

func main() {
	defer fmt.Println("This is defer #1")
	defer fmt.Println("This is defer #2")
	defer fmt.Println("This is defer #3")

	fmt.Println("main function")
}
