package main

import "fmt"

// Массив - это всегда фиксированная длинна

func main() {
	planets := [...]string{
		"Меркурий",
		"Венера",
		"Земля",
		"Марс",
		"Сатурн",
		"Уран класная планета",
	}
	for i, v := range planets {
		fmt.Println(i, v)
	}
}
