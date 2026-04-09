package main

import "fmt"

// Массив - это всегда фиксированная длинна

func change01index(arr [3]string) {
	arr[0] = "different planet"
}

func main() {
	planets := [...]string{
		"Меркурий",
		"Венера",
		"Земля",
	}

	change01index(planets)

	fmt.Println(planets)

}
