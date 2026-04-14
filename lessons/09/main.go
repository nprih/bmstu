package main

import "fmt"

type Engine struct {
	Name string
}

func (e Engine) start(arr []int) {
	fmt.Println("brbrbrbrbrbrbrb")
}

type Car struct {
	Number string
	Color  string
	Engine Engine
}

type Motorcycle struct {
	Number string
	Vendor string
	Engine Engine
}

func main() {
	myCar := Car{
		Number: "а001аа777",
		Color:  "red",
		Engine: Engine{
			Name: "B123A56",
		},
	}

	myCicle := Motorcycle{
		Number: "а0000",
		Vendor: "BMW",
		Engine: Engine{
			Name: "B123A56",
		},
	}

	fmt.Println(myCar.Engine)
	fmt.Println(myCicle.Engine)
}
