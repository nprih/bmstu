package main

import "fmt"

func main() {
	type celsius float64

	var wrongCelsius float64 = 100

	var temperature celsius = 10.7 + celsius(wrongCelsius)

	fmt.Println(temperature)
}
