package main

import (
	"fmt"
)

func main() {
	ch := make(chan int, 3) //cap = 3

	ch <- 1
	ch <- 10 //len(ch) = 2, cap(ch) = 3
	ch <- 20

	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
