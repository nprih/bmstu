package main

import (
	"fmt"
)

func main() {
	dmitry := map[string]string{
		"Surname": "Verevkin",
		"Name":    "Dmitry",
		"Email":   "veryovkin@bmstu.ru",
	}
	delete(dmitry, "Name")
	for k, v := range dmitry {
		fmt.Println(k, ":", v)
	}
}
