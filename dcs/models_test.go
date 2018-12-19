package dcs

import "testing"

func TestModelsAverage(t *testing.T) {
	amount := 1 + 1
	if amount != 2 {
		t.Errorf("0People in space, got: %d, want: %d.", amount, 1)
	}
}
