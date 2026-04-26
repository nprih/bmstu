package main

import "testing"

func InsertXInMap(x int, b *testing.B) {
	testMap := make(map[int]int, x)
	b.ResetTimer()
	for i := 0; i < x; i++ {
		testMap[i] = i
	}
}

func InsertXInSlice(x int, b *testing.B) {
	testSlice := make([]int, x)
	b.ResetTimer()
	for i := 0; i < x; i++ {
		testSlice = append(testSlice, i)
	}
}
