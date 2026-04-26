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

func BenchmarkInsertXInMapInterface100000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		InsertXInMapInterface(100000, b)
	}
}

func BenchmarkInsertXInMapInterface1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		InsertXInMapInterface(1000, b)
	}
}

func BenchmarkInsertXInMapInterfaceMap100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		InsertXInMapInterface(100, b)
	}
}
