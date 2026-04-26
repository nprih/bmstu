package main

import "testing"

func BenchmarkInsertXInMap100000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		InsertXInMap(100000, b)
	}
}

func BenchmarkInsertXInMap1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		InsertXInMap(1000, b)
	}
}

func BenchmarkInsertXInMap100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		InsertXInMap(100, b)
	}
}

func BenchmarkInsertXInSlice100000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		InsertXInSlice(100000, b)
	}
}
func BenchmarkInsertXInSlice1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		InsertXInSlice(1000, b)
	}
}
func BenchmarkInsertXInSlice100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		InsertXInSlice(100, b)
	}
}
