package main

import "fmt"

// Массив - это всегда фиксированная длина
// Слайс - не фиксированная, по умолчанию, длина. Можно увеличивать...

func main() {
	planets_2 := make([]int, 0, 10)
	for i := 0; i < 30; i++ {
		planets_2 = append(planets_2, i)
		fmt.Printf("len slice: %v,  cap slice: %v\n", len(planets_2), cap(planets_2))
	}
}
