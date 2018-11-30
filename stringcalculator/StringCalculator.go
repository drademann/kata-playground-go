// Package stringcalculator calculates the sum of numbers contained in a string.
package stringcalculator

import (
	"regexp"
	"strconv"
)

var splitter = regexp.MustCompile(`[,\n]`)

// SumOf returns the sum of the numbers contained in a string.
func SumOf(input string) int {
	if input == "" {
		return 0
	}
	return sumOf(splitted(input))
}

func splitted(input string) []string {
	return splitter.Split(input, -1)
}

func sumOf(input []string) int {
	sum := 0
	for _, s := range input {
		n, _ := strconv.Atoi(s)
		sum += n
	}
	return sum
}
