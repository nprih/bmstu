package main

import "fmt"

type Animal struct {
	Name  string
	Voice string
}

func (a Animal) talk() {
	fmt.Printf("%s говорит: %s!\n", a.Name, a.Voice)
}
func main() {
	Barsic := Animal{
		Name:  "Барсик",
		Voice: "Мяу",
	}
	Sharik := Animal{
		Name:  "Шарик",
		Voice: "Гав",
	}
	Barsic.talk()
	Sharik.talk()
}
