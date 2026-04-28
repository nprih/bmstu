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

func BenchmarkSelectXFromMap100000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SelectXFromMap(100000, b)
	}
}

func BenchmarkSelectXFromMap1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SelectXFromMap(1000, b)
	}
}

func BenchmarkSelectXFromMap100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SelectXFromMap(100, b)
	}
}

func BenchmarkSelectXFromSlice100000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SelectXFromSlice(100000, b)
	}
}
func BenchmarkSelectXFromSlice1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SelectXFromSlice(1000, b)
	}
}
func BenchmarkSelectXFromSlice100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SelectXFromSlice(100, b)
	}
}
