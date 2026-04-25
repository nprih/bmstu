package main

import "fmt"

type Book struct {
	Title  string
	Author string
	Year   int
}

var book Book

func (item Book) GetInfo() string {
	return fmt.Sprintf("\"%s\" by %s, %d", item.Title, item.Author, item.Year)
}

func main() {
	book = Book{Title: "PHP 8: объекты, шаблоны и методики программирования", Author: "Зандстра М.", Year: 2021}
	fmt.Println(book.GetInfo())
}
