package main

import (
	"fmt"
	"math/rand"
	"time"
)

func sleepyGopher(index int, ch chan int) {
	time.Sleep(time.Duration(rand.Intn(6)) * time.Second)
	ch <- index
}

// select
func main() {
	ch := make(chan int)
	for i := 0; i < 5; i++ {
		go sleepyGopher(i, ch)
	}
	for i := 0; i < 5; i++ {
		fmt.Println(<-ch)
	}
}
