package main

import "fmt"

func factorial(a uint) uint {
	var b, i uint = 1, 1
	for i <= a {
		b *= i
		i++
	}
	return b
}
func main() {
	var a uint
	fmt.Scanln(&a)
	fmt.Println(factorial(a))
}
