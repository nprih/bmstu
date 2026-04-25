package main

import "fmt"

func main() {
	runCounter()
}

func runCounter() {
	count := 0
	counter := func() int {
		count++
		return count
	}

	for i := 0; i < 3; i++ {
		fmt.Println(counter())
	}
}
