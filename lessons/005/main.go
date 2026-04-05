package main

import "fmt"

func main() {
	peace := "Hello, world"
	for i := 0; i < len(peace); i++ {
		fmt.Printf("%c\n", peace[i])
	}
}
