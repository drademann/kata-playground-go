package playground

import "testing"

func TestStringCalculator(t *testing.T) {
	var tests = []struct {
		input    string
		expected int
	}{
		{"", 0},
		{"1", 1},
		{"1,2", 3},
		{"1,2,3", 6},
		{"1\n2,3", 6},
	}
	for _, test := range tests {
		if sum := SumOf(test.input); sum != test.expected {
			t.Errorf("SumOf(%q) = %d (expected %d)", test.input, sum, test.expected)
		}
	}
}
