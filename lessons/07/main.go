package main

import "fmt"

// пустые строки

func main() {
	var planets [8]string
	planets[0] = "Меркурий"
	planets[1] = "Венера"
	planets[2] = "Земля"

	mercury := planets[0]
	earth := planets[2]

	fmt.Println(mercury, earth, planets[1])
	fmt.Println(len(planets))
	fmt.Println(planets[6] == "")

}
