package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func sleepyGopher(index int, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Duration(rand.Intn(3)) * time.Second)
	fmt.Println(index)
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 7; i++ {
		wg.Add(1)
		go sleepyGopher(i, &wg)
	}
	wg.Wait()
}
