package main

import "fmt"

func main() {
	var a int = 3
	switch {
	case a > 1:
		fmt.Println("a > 1")
		fallthrough
	case a == 100:
		fmt.Println("3")
	case a == 4 || a == 5:
		fmt.Println("4 or 5")
	default:
		fmt.Println("Default case")
	}
}
