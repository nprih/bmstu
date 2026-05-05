package main

import "fmt"

func badFunc() {
	defer func() {
		r := recover()
		if r != nil {
			fmt.Println("It was panic, but i am stronger")
			fmt.Println("panic was: ", r)
		}
	}()
	panic("I am panic Dmitry, HI")
}
func main() {
	defer func() {
		r := recover()
		if r != nil {
			fmt.Print("Panic was:", r)
		}
	}()
	var arr [2]int
	i := 2
	arr[i] = 0 //panic
}
