package main

import "fmt"

type User struct {
	Name  string
	Email string
	Age   int
}

func (u User) sendMessage(message string) {
	fmt.Printf("Message: %s was sent to %s \n", message, u.Email)
}
func main() {
	var Ivan User = User{Name: "Ivan", Email: "Ivan@mail.ru", Age: 19}
	var Kirill User = User{Name: "Kirill", Email: "Kirill@mail.ru", Age: 20}
	users := []User{Ivan, Kirill}
	for _, v := range users {
		v.sendMessage("Welcome to afterparty!")
	}
}
