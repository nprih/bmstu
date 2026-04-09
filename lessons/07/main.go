package main

import "fmt"

// Массив - это всегда фиксированная длина
// Слайс - не фиксированная, по умолчанию, длина. Можно увеличивать...

func changeZeroIndex(arr []int) {
	arr[0] = 99999
}

func main() {
	planets_2 := make([]int, 0, 100)
	for i := 0; i < 30; i++ {
		planets_2 = append(planets_2, i)
	}
	changeZeroIndex(planets_2)

	planets_2 = append(planets_2[:1], planets_2[len(planets_2)-1])

	fmt.Println(planets_2)
}
