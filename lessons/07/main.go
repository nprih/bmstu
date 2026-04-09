package main

import "fmt"

// Массив - это всегда фиксированная длина
// Слайс - не фиксированная, по умолчанию, длина. Можно увеличивать...

func main() {
	//planets := []string{}
	planets := make([]string, 100)
	fmt.Println(len(planets))
}
