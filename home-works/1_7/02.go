package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const quit = "exit"

var cities []string
var stop = ""
var newCity string
var search string

var err error
var setText = "Исходный слайс с городами: "
var addText = "Введенный город добавлен:  "
var delText = "Введенный город удален:    "

func main() {
	cities = []string{
		"Москва",
		"Санкт-Петербург",
		"Можайск",
		"Александров",
		"Одинцово",
		"Мытищи",
		"Реутов",
		"Владимир",
		"Лобня",
		"Дмитров",
	}
	fmt.Printf("%s%v\n\n", setText, cities)

	addCity()

	if stop == quit {
		return
	}

	deleteCity()

	if stop == quit {
		return
	}

	//findCity()

}

func addCity() {
	inputAddCity()
	if stop == quit {
		return
	}
	cities = append(cities, strings.TrimSpace(newCity))
	fmt.Printf("%s%v\n\n", addText, cities)
}

func deleteCity() {
	inputDeleteCity()
	if stop == quit {
		return
	}
	exist, index := find()

	if exist {
		cities = append(cities[:index], cities[index+1:]...)
		fmt.Printf("%s%v\n\n", delText, cities)
	}

}

func inputAddCity() {
	stop = ""
	fmt.Print("Ведите город, который хотите добавить: ")
	reader := bufio.NewReader(os.Stdin)
	newCity, err = reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		stop = quit
		return
	}
}

func inputDeleteCity() {
	stop = ""
	fmt.Print("Ведите город, который хотите удалить: ")
	reader := bufio.NewReader(os.Stdin)
	search, err = reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		stop = quit
		return
	}
}

func findCity() {
	inputFindCity()

	exist, index := find()
	if exist {
		fmt.Println(index)
	}
}
func inputFindCity() {
	fmt.Print("Ведите город, который хотите найти: ")
	reader := bufio.NewReader(os.Stdin)
	search, err = reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return
	}
}

func find() (bool, int) {
	mapCities := make(map[string]int)
	for i, city := range cities {
		mapCities[city] = i
	}
	if index, exist := mapCities[strings.TrimSpace(newCity)]; exist {
		return true, index
	}
	return false, 0
}
