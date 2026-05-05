package main

import "fmt"

func badFunc() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("It was panic, but i'm stronger", r)
			fmt.Println("panic was: ", r)
		}
	}()
	panic("I'm bad function")
}

func main() {
	badFunc()
}
