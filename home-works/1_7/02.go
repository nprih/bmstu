package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

var cities = []string{"Москва", "Балашиха", "Можайск", "Александров", "Одинцово",
	"Мытищи", "Реутов", "Владимир", "Лобня", "Дмитров"}
var line = strings.Repeat("-", 100)
var inpCity string
var err error
var res string

func main() {
	fmt.Printf("%s\n\n%s\n\n%s\n", line, "Список городов: ",
		getCitiesColumn(false, 0))
	addCity()
	deleteCity()
	findCity()
}

func addCity() {
	inputCity("Ведите город, который хотите добавить: ")
	cities = append(cities, strings.TrimSpace(inpCity))
	fmt.Printf("\n%s\n\nГород \"%s\" добавлен:\n\n%s\n", line, inpCity,
		getCitiesColumn(false, 0))
}

func deleteCity() {
	inputCity("Ведите город, который хотите удалить: ")
	exist, index := find()
	if exist {
		cities = append(cities[:index], cities[index+1:]...)
		res = fmt.Sprintf("\n%s\n\nВведенный город (\"%s\") удален: \n\n%s\n", line, inpCity,
			getCitiesColumn(false, 0))
	} else {
		res = fmt.Sprintf("\n%s\n\nГород \"%s\" не найден в списке\n\n%s\n\n", line, inpCity, line)
	}
	fmt.Printf(res)
}

func findCity() {
	inputCity("Ведите город, который хотите найти: ")
	exist, index := find()
	if exist {
		res = fmt.Sprintf("\n%s\n\nВыбранный город (\"%s\") выделен в списке:\n\n%s\n", line, inpCity,
			getCitiesColumn(exist, index))
	} else {
		res = fmt.Sprintf("\n%s\n\nГород \"%s\" не найден в списке\n\n%s\n\n", line, inpCity, line)
	}
	fmt.Printf(res)
}

func inputCity(message string) {
	fmt.Print(message)
	reader := bufio.NewReader(os.Stdin)
	inpCity, err = reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return
	}
	capitalizeWords()
}

func getCitiesColumn(selected bool, index int) (res string) {
	for i, city := range cities {
		if selected && index == i {
			res += "~ " + city + " ~" + "\n"
		} else {
			res += city + "\n"
		}
	}
	return res
}

func find() (bool, int) {
	mapCities := make(map[string]int)
	for i, city := range cities {
		mapCities[city] = i
	}
	if index, exist := mapCities[inpCity]; exist {
		return true, index
	}
	return false, 0
}

func capitalizeWords() {
	inpCity = strings.TrimSpace(inpCity)
	char := []rune(inpCity)
	for i := range char {
		if i == 0 {
			char[i] = unicode.ToUpper(char[i])
		} else {
			char[i] = unicode.ToLower(char[i])
		}
	}
	inpCity = string(char)
}
