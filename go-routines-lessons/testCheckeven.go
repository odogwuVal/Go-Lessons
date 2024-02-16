package main

import "testing"

func TestCheckEven(t *testing.T) {
	i := 10
	expected := "YES"
	res := CheckEven(i)

	if res != expected {
		t.Errorf("Expected: %s, got: %s", expected, res)
	}
}
