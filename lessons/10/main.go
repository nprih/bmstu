package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	One   = "Вывести весь \"арсенал\" библиотеки (вообще все книги в библиотеке) | красиво!!"
	Two   = "Вывести только доступные/недоступные книги (выбор от пользователя)"
	Three = "Вывести только книги, отсортированные по автору (автора вводит пользователь)"
	Four  = "Взять книгу себе (изменить флаг) | предусмотреть невозможность взять недоступную книгу"
	Five  = "5 - Выход из приложения\n"
	quit  = "quit"
)

type Book struct {
	Title     string
	Author    string
	Available bool
}

var numTask int
var stop string
var books = []Book{
	{
		Title:     "Мёртвые души",
		Author:    "Гоголь",
		Available: true,
	},
	{
		Title:     "Война и мир",
		Author:    "Толстой",
		Available: true,
	},
	{
		Title:     "Капитанская дочка",
		Author:    "Пушкин",
		Available: true,
	},
	{
		Title:     "Сказка о Царе Салтане",
		Author:    "Пушкин",
		Available: true,
	},
}

// by Николай
func main() {
	for stop != quit {
		numTask = 0
		taskSelection()
		taskSolution()
	}
}

// by Николай
func taskSolution() {
	switch numTask {
	case 1:
		printAllBook(books)
	case 2:
		filterByAvailability(books)
	case 3:
		filterBooksByAuthor(books)
	case 4:
		takeBook(books)
	}
}

// by Николай
func taskSelection() {
	fmt.Printf("\n1 - %s\n2 - %s\n3 - %s\n4 - %s\n\n%s\nВведите номер задачи: ", One, Two, Three, Four, Five)
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return
	}
	checkNumber(text)
}

// by Николай
func checkNumber(text string) {
	numberTask, err := strconv.Atoi(strings.TrimSpace(text))
	switch true {
	case err != nil:
		fmt.Print(strings.TrimSpace(text), " - не корректное число\n")
		return
	case numberTask <= 0:
		fmt.Print(strings.TrimSpace(text), " - не корректное число\n")
		return
	case numberTask == 5:
		fmt.Print("Выход из приложения", "\n")
		stop = quit
		return
	case numberTask > 5:
		fmt.Print("У нас не так много задач\n")
		return
	default:
		numTask = numberTask
		fmt.Print("Выбрана задача: ", numTask, "\n")
	}
}

// by Арсений
func printAllBook(books []Book) {
	printHeader()
	for i, book := range books {
		printBook(i, book)
	}
	printFooter()
}

// by Арсений
func filterByAvailability(books []Book) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Показать (1 - доступные, 2 - недоступные): ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	var flag bool
	if input == "1" {
		flag = true
	} else if input == "2" {
		flag = false
	} else {
		fmt.Println("Неверный ввод")
		return
	}

	fmt.Println("\nРезультат:")
	printHeader()
	for i, b := range books {
		if b.Available == flag {
			printBook(i, b)
		}
	}
	printFooter()
}

// by Арсений
func printBook(i int, book Book) {
	status := "Недоступна"
	if book.Available {
		status = "Доступна"
	}
	fmt.Printf("%-4d | %-30s | %-30s | %-12s\n", i, book.Title, book.Author, status)
}

// by Арсений
func printHeader() {
	fmt.Printf("%s%s%s\n", strings.Repeat("-", 39), "Каталог книг", strings.Repeat("-", 39))
	fmt.Printf("%-4s | %-30s | %-30s | %-12s \n", "№", "Название", "Автор", "Доступность")
	fmt.Println(strings.Repeat("-", 90))
}

// by Арсений
func printFooter() {
	fmt.Println(strings.Repeat("-", 90))
}

// by Анатолий
func filterBooksByAuthor(books []Book) {
	var filteredBooks []Book
	var author string
	fmt.Println("Введите фамилию автора:")
	fmt.Scan(&author)
	for _, book := range books {
		if book.Author == author {
			filteredBooks = append(filteredBooks, book)
		}
	}
	printAllBook(filteredBooks)
}

// by Анатолий
func takeBook(books []Book) bool {
	printAllBook(books)
	var userInput string
	fmt.Println("Введите идентификатор книги:")
	fmt.Scan(&userInput)
	bookIndex, err := strconv.Atoi(userInput)
	if err != nil {
		fmt.Println("Вы ввели некорректный идентификатор")
		return false
	}
	if bookIndex >= len(books) || bookIndex < 0 {
		fmt.Println("Книги с таким идентификатором нет в нашей библиотеке")
		return false
	}
	if books[bookIndex].Available {
		books[bookIndex].Available = false
		fmt.Printf("Вы успешно взяли книгу %s\n", books[bookIndex].Title)
		return true
	}
	fmt.Printf("Книгу %s уже взяли\n", books[bookIndex].Title)
	return false
}
