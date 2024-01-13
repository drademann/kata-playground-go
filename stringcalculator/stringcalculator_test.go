package stringcalculator

import "testing"

func TestAdd_givenEmptyString_shouldReturnZero(t *testing.T) {
	result := Add("")

	if result != 0 {
		t.Errorf("Add(\"\") = %d, want 0", result)
	}
}

func TestAdd_givenSingleNumber_shouldReturnIt(t *testing.T) {
	result := Add("5")

	if result != 5 {
		t.Errorf("Add(\"5\") = %d, want 5", result)
	}
}
