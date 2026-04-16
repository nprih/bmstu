package main

import "fmt"

type Cat struct {
	Name  string
	Voice string
}

type Dog struct {
	Name  string
	Voice string
}

func (c Cat) talk() {
	fmt.Printf("Cat tolk %s\n", c.Voice)
}

func (c Dog) talk() {
	fmt.Printf("Dog tolk %s\n", c.Voice)
}

func main() {
	Barsik := Cat{Name: "Barsik", Voice: "Meuw"}
	Sharik := Dog{Name: "Sharik", Voice: "Wuf!"}

	Barsik.talk()
	Sharik.talk()

}
