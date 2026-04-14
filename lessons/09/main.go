package main

import "fmt"

type User struct {
	Name  string
	Email string
	Age   int
}

func main() {
	var Ivan User = User{Name: "Ivan", Email: "Ivan@email.ru", Age: 19}
	var Kirill User = User{Name: "Kirill", Email: "Kirill@email.ru", Age: 20}

	users := []User{Ivan, Kirill}

	fmt.Println(users)
}
