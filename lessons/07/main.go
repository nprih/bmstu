package main

import "fmt"

// Массив - это всегда фиксированная длинна

func main() {
	planets := [...]string{
		"Меркуий",
		"Венера",
		"Земля",
		"Марс",
		"Юпитер",
		"Сатурн",
		"Уран",
	}
	allPlanets := planets[:]

	fmt.Println(allPlanets)
}
