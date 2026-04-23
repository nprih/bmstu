package main

import (
	"fmt"
	"os"
)

func main() {
	flag := os.O_RDWR | os.O_CREATE | os.O_APPEND
	file, err := os.OpenFile("test.txt", flag, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	text := "\nffrf"
	_, err = file.WriteString(text)
	if err != nil {
		fmt.Println(err)
		return
	}
}
