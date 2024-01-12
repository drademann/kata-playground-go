package stringcalculator

import "testing"

/*
 	The String Calculator Kata

	Provide the following method

		int Add(string numbers)

	Requirements:

		1. The method returns 0 given an empty string.
		2. The method returns the number given the string only contains a single number.
		3. Given the string contains two comma-separated numbers, then the method returns their sum.
*/

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
