package main

import "fmt"

func main() {
	a := 1.0 / 3
	fmt.Println(a)
	fmt.Printf("%v\n", a)
	fmt.Printf("%f\n", a)
	fmt.Printf("%.3f\n", a)
	fmt.Printf("%07.2f\n", a)
}
