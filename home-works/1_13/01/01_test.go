package main

import (
	"errors"
	"testing"
)

type divideOutput struct {
	result float64
	err    error
}

func TestDivide(t *testing.T) {
	var testCases = []struct {
		description string
		input       [2]float64
		want        divideOutput
	}{
		{"Положительный тест деления 128 / 8",
			[2]float64{128, 8}, divideOutput{16, nil}},
		{"Положительный тест деления 256 / 16", [2]float64{256, 16},
			divideOutput{16, nil}},
		{"Положительный тест деления 512 / 64", [2]float64{512, 64},
			divideOutput{8, nil}},
		{"Отрицательный тест деления 1024 / 0", [2]float64{1024, 0},
			divideOutput{0, DivideByZero}},
	}

	for _, tt := range testCases {
		t.Run(tt.description, func(t *testing.T) {
			res, err := Divide(tt.input[0], tt.input[1])
			if res != tt.want.result || !errors.Is(err, tt.want.err) {
				t.Errorf("Тест провален!!! Ожидалось: (%.2f, %s); "+
					"Получено (%.2f, %s)", tt.want.result, tt.want.err, res, err)
			}
		})
	}
}
