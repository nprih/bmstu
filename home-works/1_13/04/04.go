package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

func main() {
	for _, slice := range GetRandomArg() {
		fmt.Printf("Числа: %v; Сумма чисел: %d\n", SliceToString(slice), SumAll(slice...))
	}
}

func GetRandomArg() (slices [][]int) {
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

func SliceToString(slice []int) string {
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

func SumAll(nums ...int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}
