package main

import "fmt"

func CheckAge(age int, name string) {
	if age >= 18 {
		fmt.Printf("Можете проходить в наш бар, %s \n", name)
	} else if age <= 0 {
		fmt.Println("Вы врете")
	} else {
		fmt.Println("Тебе пока рано")
	}
}

func main() {
	var name string
	var age int

	fmt.Print("Введите имя: ")
	fmt.Scanln(&name)

	fmt.Print("Введите возраст: ")
	fmt.Scanln(&age)

	CheckAge(age, name)
}
