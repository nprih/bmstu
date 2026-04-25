package main

import (
	"fmt"
)

type Student struct {
	Name   string
	Grades []float64
}

func (student Student) AverageGrade() float64 {
	sum := 0.0
	for _, grade := range student.Grades {
		sum += grade
	}
	return sum / float64(len(student.Grades))
}

func main() {
	student := Student{Name: "Алёша", Grades: []float64{3.0, 4.0, 4.0}}
	fmt.Printf("Средний бал студента с именем %s: %.2f\n", student.Name, student.AverageGrade())
}
