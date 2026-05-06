package main

import (
	"fmt"
	"sync"
)

const countNum = 10

func generator(ch chan int, wg *sync.WaitGroup) {
	for i := 1; i <= countNum; i++ {
		ch <- i
	}
	close(ch)
	wg.Wait()
	fmt.Printf("\nКанал закрыт.\n\n")
}

func printer(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		num, opened := <-ch
		if opened == false {
			return
		}
		fmt.Printf("Прочитано из канала: %d\n", num)
	}
}

func run(ch chan int, wg *sync.WaitGroup) {
	go generator(ch, wg)
	go printer(ch, wg)
}

func main() {
	ch := make(chan int)
	var wg sync.WaitGroup
	wg.Add(1)
	run(ch, &wg)
	wg.Wait()
	fmt.Println("Выполнение программы завершено.")
}
