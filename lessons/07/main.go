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
	FirstPlanets := planets[0:3]
	fmt.Println(FirstPlanets)

	SecondPlanets := planets[3:6]
	fmt.Println(SecondPlanets)

	FirstPlanets[0] = "different planet"
	SecondPlanets[0] = "different planet"
	fmt.Println(planets)
}
