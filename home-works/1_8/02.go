package main

import "fmt"

func main() {
	a, b := 100, 7
	printResult(a, b)
}

func printResult(a int, b int) {
	result := ""
	result += fmt.Sprintf("%d + %d = %d\n", a, b, applyOperation(a, b, add))
	result += fmt.Sprintf("%d - %d = %d\n", a, b, applyOperation(a, b, sub))
	result += fmt.Sprintf("%d * %d = %d\n", a, b, applyOperation(a, b, mul))
	if b == 0 {
		result += fmt.Sprintf("%d / %d - На ноль делить нельзя, так сказал калькулятор!\n", a, b)
	} else {
		result += fmt.Sprintf("%d / %d = %d\n", a, b, applyOperation(a, b, div))
	}
	fmt.Println(result)
}

func applyOperation(a, b int, op func(int, int) int) int {
	return op(a, b)
}

func add(x, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}

func mul(x, y int) int {
	return x * y
}

func div(x, y int) int {
	if y == 0 {
		return 0
	}
	return x / y
}
