package main

import "fmt"

const costA = 12.5
const costB = 8.3

func calc(a, b uint) (float64, float64) {
	sumA := float64(a) * costA
	sumB := float64(b) * costB

	return sumA, sumB
}

func main() {
	var a, b uint

	fmt.Printf("Введите количество километров: ")
	fmt.Scanln(&a)
	fmt.Printf("Введите количество минут: ")
	fmt.Scanln(&b)

	sumA, sumB := calc(a, b)

	fmt.Printf("\nЦена поездки в км: %.2f\n", sumA)
	fmt.Printf("Цена поездки в минутах: %.2f\n", sumB)
	res := ""
	if sumA > sumB {
		res = "по минутам"
	} else {
		res = "по километрам"
	}
	fmt.Printf("\nЛучше использовать тариф: %s\n", res)
}
