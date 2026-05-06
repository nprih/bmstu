package main

import (
	"fmt"
	"sync"
	"time"
)

const countGoRoutine = 3
const sleep = 1

var count int

func sleepyGopher(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Горутина %d: начала работать\n", id)

	count++
	if count == countGoRoutine {
		fmt.Printf("\nОжидаем завершения работы горутин...\n\n")
	}

	time.Sleep(time.Duration(sleep) * time.Second)
	fmt.Printf("Горутина %d: завершила работу\n", id)
}

func main() {
	var wg sync.WaitGroup
	for i := 1; i <= countGoRoutine; i++ {
		wg.Add(1)
		go sleepyGopher(i, &wg)
	}
	wg.Wait()
	fmt.Printf("\nВсе горутины закончили работу!\n")
}
