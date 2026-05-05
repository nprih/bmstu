package main

import "fmt"

func main() {
	defer func() {
		r := recover()
		if r != nil {
			fmt.Println("Panic was:", r)
		}
	}()
	var arr [2]int
	i := 2
	arr[i] = 0 //panic
}
