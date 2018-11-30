// Package primefactors determines the prime factors of an integer number.
package primefactors

func PrimeFactors(number int) []int {
	var factors []int
	candidate := 2
	for number > 1 {
		for ; number%candidate == 0; number /= candidate {
			factors = append(factors, candidate)
		}
		candidate++
	}
	return factors
}
