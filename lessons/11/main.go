package main

import (
	"fmt"
	"strings"
)

func providerGopher(ch12 chan string) {
	data := []string{"apple", "pen", "", "pineapple", "bad", "bad apple", "hello world"}
	for _, v := range data {
		ch12 <- v
	}
	close(ch12)
}
func analyzeGopher(ch12, ch23 chan string) {
	//defer func() {
	//	if err := recover(); err != nil {
	//		fmt.Println("I fixed.")
	//		close(ch23)
	//	}
	//}()
	for {
		item, ok := <-ch12
		if !ok {
			//ch12 <- "asdasd" //ошибка
			close(ch23)
			return
		}
		if !strings.Contains(item, "bad") {
			ch23 <- item
		}
	}
}
func printerGopher(ch23 chan string) {
	for {
		item, ok := <-ch23
		if !ok {
			return
		}
		fmt.Printf("%s, ", item)
	}
}
func main() {
	ch12 := make(chan string)
	ch23 := make(chan string)
	go analyzeGopher(ch12, ch23)
	go providerGopher(ch12)
	printerGopher(ch23)
}
