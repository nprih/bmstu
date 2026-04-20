package main

import (
	"fmt"
	"time"
)

func sleepyGopher(index int, c chan int) {
	time.Sleep(3 * time.Second)
	c <- index
}

func main() {
	c := make(chan int)

	for i := 0; i < 5; i++ {
		go sleepyGopher(i, c)
	}
	for i := 0; i < 5; i++ {
		gopherId := <-c
		fmt.Printf("Gopher %d finished work\n", gopherId)
	}
}
