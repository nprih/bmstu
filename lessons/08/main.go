package main

import "fmt"

func Sum(x ...int) int {
	res := 0
	for _, v := range x {
		res += v
	}
	return res
}

func main() {
	fmt.Println(Sum(5, 2, 3))
}
