package main

import "fmt"

func main() {
	dmitry := map[string]string{
		"Surname": "Verevkin",
		"Name":    "Dmitry",
		"Email":   "veryovkin@bmstu.ru",
	}
	if dmitry["Age"] != "" {
		fmt.Println(dmitry["Age"])
	} else {
		fmt.Println("I havo no age, sorry =( ")
	}
}
