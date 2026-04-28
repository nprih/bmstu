package main

import (
	"fmt"
	"math"
)

const circle = "окружность"
const rectangle = "прямоугольник"

type Shape interface {
	GetName() string
	Area() float64
}

type Circle struct {
	name   string
	radius float64
}

func (c Circle) Area() float64 {
	return c.radius * c.radius * math.Pi
}

func (c Circle) GetName() string {
	return circle
}

type Rectangle struct {
	name   string
	width  float64
	height float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func (r Rectangle) GetName() string {
	return rectangle
}

func main() {
	shapes := []Shape{
		&Circle{radius: 3},
		&Rectangle{width: 4, height: 5},
		&Circle{radius: 5},
		&Rectangle{width: 6, height: 7},
	}
	printAreas(shapes)
}

func printAreas(shapes []Shape) {
	for i, shape := range shapes {
		fmt.Printf("Фигура %d (%s): площадь = %.2f\n", i+1, shape.GetName(), shape.Area())
	}
}
