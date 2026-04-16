package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

var sliceCities []string
var indexCity int
var newCity string
var serCity string
var finCity string
var err error
var setText = "Исходный слайс с городами: "
var addText = "Введенный город добавлен:  "
var delText = "Введенный город удален:    "

func main() {
	sliceCities = []string{
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
	fmt.Printf("%s%v\n", setText, sliceCities)

	inputAddCity()

	sliceCities = append(sliceCities, strings.TrimSpace(newCity))
	sort.Strings(sliceCities)

	fmt.Printf("%s%v\n", addText, sliceCities)

	inputDeleteCity()
	deleteCity()
	fmt.Printf("%s%v\n", delText, sliceCities)

	inputFindCity()
	fmt.Println(finCity)

}

func inputAddCity() {
	fmt.Print("Ведите город, который хотите добавить: ")
	reader := bufio.NewReader(os.Stdin)
	newCity, err = reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return
	}
}

func inputDeleteCity() {
	fmt.Print("Ведите город, который хотите удалить: ")
	reader := bufio.NewReader(os.Stdin)
	serCity, err = reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return
	}
}

func inputFindCity() {
	fmt.Print("Ведите город, который хотите удалить: ")
	reader := bufio.NewReader(os.Stdin)
	finCity, err = reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return
	}
}

func deleteCity() {
	fmt.Println(serCity)
	findCity()
	sliceCities = append(sliceCities[:indexCity], sliceCities[indexCity+1:]...)
}

func findCity() {
	indexCity = sort.SearchStrings(sliceCities, strings.TrimSpace(serCity))
	if indexCity < len(sliceCities) {
		if sliceCities[indexCity] != strings.TrimSpace(serCity) {

			fmt.Println("Такой город не найден в данном слайсе")
			return
		}
	} else {
		fmt.Println("Такой город не найден в данном слайсе")
		return
	}
	fmt.Println(indexCity)
}
