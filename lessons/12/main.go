package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.OpenFile("test.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	text := "\nvvvvvvvvv"
	_, err = file.WriteString(text)
	if err != nil {
		fmt.Println(err)
		return
	}
}
