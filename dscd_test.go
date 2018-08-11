package main

import "testing"

func TestAverage(t *testing.T) {
	amount := 1 + 1
	if amount != 1 {
		t.Errorf("People in space, got: %d, want: %d.", amount, 1)
	}
}
