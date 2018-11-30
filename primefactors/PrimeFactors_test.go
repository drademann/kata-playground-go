// Requirement: For an integer value n the generator should return all prime factors as an ordered list.
package primefactors

import "testing"

func TestPrimeFactors(t *testing.T) {
	var tests = []struct {
		input    int
		expected []int
	}{
		{1, []int{}},
		{2, []int{2}},
		{3, []int{3}},
		{4, []int{2, 2}},
		{6, []int{2, 3}},
		{9, []int{3, 3}},
	}
	for _, test := range tests {
		if factors := PrimeFactors(test.input); notEqual(factors, test.expected) {
			t.Errorf("PrimeFactors(%d) = %d (expected %d)", test.input, factors, test.expected)
		}
	}
}

func notEqual(ints1, ints2 []int) bool {
	if len(ints1) != len(ints2) {
		return true
	}
	for i, i1 := range ints1 {
		if i1 != ints2[i] {
			return true
		}
	}
	return false
}
