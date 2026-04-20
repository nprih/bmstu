package main

import (
	"fmt"
	"time"
)

func sleepyGopher(index int) {
	time.Sleep(3 * time.Second)
	fmt.Println("hrrr... #", index)
}

func main() {
	for i := 0; i < 5; i++ {
		go sleepyGopher(i)
	}
	time.Sleep(4 * time.Second)
}
