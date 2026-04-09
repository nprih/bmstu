package main

import "fmt"

// Массив - это всегда фиксированная длина
// Слайс - не фиксированная, по умолчанию, длина. Можно увеличивать...

func main() {
	planets_2 := []string{"Москва", "Сочи", "Лондон", "Череповец"}
	fmt.Println(len(planets_2))
	planets_2 = append(planets_2[:2], planets_2[3:]...)
	fmt.Println(planets_2)
	fmt.Println(len(planets_2))
}
