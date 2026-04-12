package main

import (
	"fmt"
)

func main() {
	//dmitry := make(map[string]string)
	dmitry := map[string]string{
		"Surname": "Verevkin",
		"Name":    "Dmitry",
		"Email":   "veryovkin@bmstu.ru",
	}
	for k, v := range dmitry {
		fmt.Println(k, ":", v)
	}
}
