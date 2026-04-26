package main

import "testing"

func InsertXInMap(x int, b *testing.B) {
	testMap := make(map[int]int, 0)
	b.ResetTimer()
	for i := 0; i < x; i++ {
		testMap[i] = i
	}
}
