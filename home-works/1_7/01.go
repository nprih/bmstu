package main

import (
	"fmt"
	"math/rand"
	"slices"
)

var rndText = "Рандомно заполненный массив: "
var srtText = "Отсортированный слайс:       "

func main() {
	randomArr := getRandomArr()
	fmt.Printf("%s%v\n%s%v\n", rndText, randomArr, srtText, getSortedSlice(randomArr))
}

func getRandomArr() [10]int {
	var arr [10]int
	for i, _ := range arr {
		arr[i] = rand.Intn(100) + 1
	}
	return arr
}

func getSortedSlice(randomArr [10]int) []int {
	var slice []int
	for _, v := range randomArr {
		slice = append(slice, v)
	}
	slices.Sort(slice)
	return slice
}
