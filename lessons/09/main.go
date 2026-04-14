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
	Engine
}

func main() {
	myCar := Car{
		Number: "а001аа777",
		Color:  "red",
		Engine: Engine{
			Name: "B123A56",
		},
	}

	myCar.start()
}

//type Engine struct {
//	Name string
//}
//func (e Engine) start() {
//	fmt.Println("brbrbrbrbrbrbr")
//}
//type Car struct {
//	Number string
//	Color  string
//	Engine
//}
//func main() {
//	myCar := Car{
//		Number: "a001aa777",
//		Color:  "red",
//		Engine: Engine{
//			Name: "B123A56",
//		},
//	}
//	myCar.start()
//}
