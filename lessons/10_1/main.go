package main

import "fmt"

type addressPrinter interface{} //any

type Cat struct {
	Name  string
	Voice string
}

func (c Cat) talk() {
	fmt.Printf("Cat tolk %s\n", c.Voice)
}

func (c Cat) printAddr() string {
	return c.Name + " myCat"
}

type Dog struct {
	Name  string
	Voice string
}

func (d Dog) talk() {
	fmt.Printf("Dog tolk %s\n", d.Voice)
}
func (d Dog) printAddr() string {
	return d.Name + " myDog"
}
func PrintAddressInConsole(animal addressPrinter) {
	fmt.Println(animal)
}

func main() {
	Barsik := Cat{Name: "Barsic", Voice: "Muew!"}
	Sharik := Dog{Name: "Sahrik", Voice: "Wuf!"}
	PrintAddressInConsole(Barsik)
	PrintAddressInConsole(Sharik)
}
