package main

import (
	"fmt"
	"sort"
)

func getResult(arr []int) string {
	sorted, mean := calc(arr)
	res := "\nвыход:\n\"Отсортированные элементы: "
	for i := 0; i < len(sorted); i++ {
		res += fmt.Sprintf("%d ", sorted[i])
	}
	res = res[:len(res)-1]
	return res + fmt.Sprintf("\nСамое большое число: %d\n", sorted[0]) +
		fmt.Sprintf("Самое маленькое число: %d\n", sorted[4]) +
		fmt.Sprintf("Среднее арифметическое: %d\n\"\n", mean)
}

func calc(arr []int) ([]int, int) {
	items, summ := len(arr), 0
	for i := 0; i < len(arr); i++ {
		summ += arr[i]
	}
	sort.Sort(sort.Reverse(sort.IntSlice(arr)))
	return arr, summ / items
}

func main() {
	var a, b, c, d, e int
	fmt.Printf("вход: ")
	fmt.Scanln(&a, &b, &c, &d, &e)
	arr := []int{a, b, c, d, e}
	fmt.Println(getResult(arr))
}
