package main

import "testing"

func TestIsEven(t *testing.T) {
	var testsCases = []struct {
		description string
		input       int
		want        string
	}{
		{"Positive test on 2", 2, "Yes"},
		{"Positive test on 3", 3, "No"},
		{"Positive test on 0", 0, "Yes"},
	}
	for _, tt := range testsCases {
		t.Run(tt.description, func(t *testing.T) {
			result := IsEven(tt.input)
			if result != tt.want {
				t.Errorf("Incorrect test, want %s, but got %s ", result, tt.want)
			}
		})
	}
}
