package main

import "fmt"

func mult(a, b int) int {
	return a * b
}

func main() {
	var a, b int = 2, 3
	fmt.Printf("%d * %d = %d\n", a, b, mult(a, b))
}
