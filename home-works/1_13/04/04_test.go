package main

import "testing"

func TestSumAll(t *testing.T) {
	var testCases = []struct {
		description string
		input       []int
		want        int
	}{
		{"Положительный тест для 8, 16, 32, 64, 128, 256, 512",
			[]int{8, 16, 32, 64, 128, 256, 512}, 1016},
		{"Положительный тест для 8 / 3, 16, -32, 64, -128, 256, 512",
			[]int{8 / 3, 16, -32, 64, -128, 256, 512}, 690},
	}
	for _, tt := range testCases {
		t.Run(tt.description, func(t *testing.T) {
			result := SumAll(tt.input...)
			if result != tt.want {
				t.Errorf("Тест провален!!! Ожидалось: %d; Получено: %d", tt.want, result)
			}
		})
	}
}
