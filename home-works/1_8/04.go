package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

func main() {
	for _, slice := range getRandomArg() {
		fmt.Printf("Числа: %v; Сумма чисел: %d\n", sliceToString(slice), sumAll(slice...))
	}
}

func getRandomArg() (slices [][]int) {
	var nums []int
	countItem := rand.Intn(5) + 1
	for item := 0; item < countItem; item++ {
		nums = nums[:0]
		countInt := rand.Intn(5) + 1
		for i := 0; i <= countInt; i++ {
			nums = append(nums, rand.Intn(10)+1)
		}
		slices = append(slices, nums)
	}
	return slices
}

func sliceToString(slice []int) string {
	text := ""
	for i := range slice {
		number := slice[i]
		if i < len(slice)-1 {
			text += strconv.Itoa(number) + ", "
		} else {
			text += strconv.Itoa(number)
		}
	}
	return text
}

func sumAll(nums ...int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}
