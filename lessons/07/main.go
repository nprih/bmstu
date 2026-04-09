package main

import "fmt"

// Массив - это всегда фиксированная длина
// Слайс - не фиксированная, по умолчанию, длина. Можно увеличивать...

func main() {
	planets_2 := []int{}
	for i := 0; i < 30; i++ {
		planets_2 = append(planets_2, i)
	}
	fmt.Println(cap(planets_2))
}
