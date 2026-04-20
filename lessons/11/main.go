package main

import (
	"fmt"
	"strings"
	"time"
)

func providerGopher(ch12 chan string) {
	data := []string{"apple", "pen", "pineapple", "bad", "bad apple", "hello world"}
	for _, v := range data {
		time.Sleep(1 * time.Second)
		ch12 <- v
	}
	ch12 <- ""
}

func analyzeGopher(ch12, ch23 chan string) {
	for {
		item := <-ch12
		if item == "" {
			ch23 <- ""
			return
		}
		if !strings.Contains(item, "bad") {
			ch23 <- item
		}
	}
}

func printerGopher(ch23 chan string) {
	for {
		item := <-ch23
		if item == "" {
			return
		}
		fmt.Printf("%s, ", item)
	}
}

func main() {
	ch12 := make(chan string)
	ch23 := make(chan string)

	go providerGopher(ch12)
	go analyzeGopher(ch12, ch23)

	printerGopher(ch23)
}
