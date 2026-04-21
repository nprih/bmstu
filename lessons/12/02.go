package main

import (
	"fmt"
	"sync"
)

func main() {
	shared := make(map[int]int)

	var mu sync.Mutex
	var wg sync.WaitGroup

	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func(index int) {
			mu.Lock()
			shared[index] = index
			mu.Unlock()

			wg.Done()
		}(i)
	}

	wg.Wait()
	fmt.Println("Map lenth:", len(shared))
}
