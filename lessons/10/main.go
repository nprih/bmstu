package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Book struct {
	Title     string
	Author    string
	Available bool
}

const One = "Вывести весь \"арсенал\" библиотеки (вообще все книги в библиотеке) | красиво!!"
const Two = "Вывести только доступные/недоступные книги (выбор от пользователя)"
const Three = "Вывести только книги, отсортированные по автору (автора вводит пользователь)"
const Four = "Взять книгу себе (изменить флаг) | предусмотреть невозможность взять недоступную книгу"
const Five = "5 - Выход из приложения\n"
const endTask = "\n\nЗадача выполнена\n\n"

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

const quit = "quit"

func main() {
	for stop != quit {
		numTask = 0
		taskSelection()
		taskSolution()
	}
}

func taskSolution() {
	switch numTask {
	case 1:
		printAllBook(books)
	case 2:
		filterByAvailability(books)
	case 3:
		filterBooksByAuthor(books)
	case 4:
		//books.take()
	}
}

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

func printAllBook(books []Book) {
	printHeader()
	for i, book := range books {
		printBook(i, book)
	}
	printFooter()
}

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

func printBook(i int, book Book) {
	status := "Недоступна"
	if book.Available {
		status = "Доступна"
	}
	fmt.Printf("%-4d | %-30s | %-30s | %-12s\n", i, book.Title, book.Author, status)
}

func printHeader() {
	fmt.Printf("%s%s%s\n", strings.Repeat("-", 39), "Каталог книг", strings.Repeat("-", 39))
	fmt.Printf("%-4s | %-30s | %-30s | %-12s \n", "№", "Название", "Автор", "Доступность")
	fmt.Println(strings.Repeat("-", 90))
}
func printFooter() {
	fmt.Println(strings.Repeat("-", 90))
	fmt.Printf(endTask)
}

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
	fmt.Printf(endTask)
}

func (b *Book) take() bool {
	if b.Available {
		b.Available = false
		fmt.Printf("Вы успешно взяли книгу %s\n", b.Title)
		return true
	}
	fmt.Printf("Книгу %s уже взяли\n", b.Title)
	return false
}
