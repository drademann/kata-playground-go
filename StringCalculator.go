package playground

import "strconv"

// SumOf returns the sum of the numbers contained in a string.
func SumOf(input string) int {
	if input == "" {
		return 0
	}
	i, _ := strconv.Atoi(input)
	return i
}
