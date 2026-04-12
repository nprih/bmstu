package main

import "fmt"

func main() {
	fmt.Println(Cube(5))
}

func Cube(x int) (y int) {
	y = x * x * x
	return
}
