package main

import "fmt"

func Sum(x ...int) (res, zero int) {
	for _, v := range x {
		res += v
	}
	return res, zero
}

func main() {
	fmt.Println(Sum(5, 2, 3))
}
