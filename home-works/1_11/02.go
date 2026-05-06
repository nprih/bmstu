package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	ch := make(chan int)

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		for i := 1; i <= 10; i++ {
			ch <- i
		}
		close(ch)
		fmt.Println("Канал закрыт")
	}()

	go func() {
		defer wg.Done()
		time.Sleep(time.Duration(30) * time.Second)
		for {
			num, opened := <-ch
			if opened == false {
				return
			}
			fmt.Println("Получено из канала: ", num)
		}
	}()

	wg.Wait()
}
