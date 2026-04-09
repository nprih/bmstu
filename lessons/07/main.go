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
	for i, _ := range planets {
		fmt.Println(planets[i])
		planets[i] = "Bad planet"
	}
	fmt.Println(planets)
}
