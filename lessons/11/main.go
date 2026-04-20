package main

import (
	"fmt"
	"strings"
	"time"
)

func providerGopher(ch12 chan string) {
	data := []string{"apple", "pen", "", "pineapple", "bad", "bad apple", "hello world"}
	for _, v := range data {
		time.Sleep(1 * time.Second)
		ch12 <- v
	}
	close(ch12)
}

func analyzeGopher(ch12, ch23 chan string) {
	for item := range ch12 {
		if !strings.Contains(item, "bad") {
			ch23 <- item
		}
	}
	close(ch23)
}

func printerGopher(ch23 chan string) {
	for item := range ch23 {
		fmt.Printf("%s, ", item)
	}
}
func main() {
	ch12 := make(chan string)
	ch23 := make(chan string)

	go analyzeGopher(ch12, ch23)
	go providerGopher(ch12)

	printerGopher(ch23)
	fmt.Println()
}
