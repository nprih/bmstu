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

func SelectXFromMap(x int, b *testing.B) {
	testMap := make(map[int]int, x)
	for i := 0; i < x; i++ {
		testMap[i] = i
	}
	var h int
	b.ResetTimer()
	for i := 0; i < x; i++ {
		h = testMap[i]
	}
	if h != x {
	}
}

func SelectXFromSlice(x int, b *testing.B) {
	testSlice := make([]int, x)
	for i := 0; i < x; i++ {
		testSlice[i] = i
	}
	var h int
	b.ResetTimer()
	for i := 0; i < x; i++ {
		h = testSlice[i]
	}
	if h != x {
	}
}
