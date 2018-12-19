package main

import "testing"

func TestAverage(t *testing.T) {
	amount := 1 + 2
	if amount != 3 {
		t.Errorf("People in space, got: %d, want: %d.", amount, 1)
	}
}
