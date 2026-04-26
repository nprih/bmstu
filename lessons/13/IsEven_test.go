package main

import "testing"

func TestIsEven(t *testing.T) {
	result1 := IsEven(2)
	if result1 != "Yes" {
		t.Errorf("Incorrect resilt, got: %s, but want: %s", result1, "Yes")
	}
	t.Log("Test with param 2 ok")

	result2 := IsEven(3)
	if result2 != "No" {
		t.Errorf("Incorrect resilt, got: %s, but want: %s", result2, "No")
	} else {
		t.Log("Test with param 3 fail")
	}

}
