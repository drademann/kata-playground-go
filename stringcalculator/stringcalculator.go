package stringcalculator

import "strconv"

func Add(numbers string) int {
	if i, err := strconv.Atoi(numbers); err == nil {
		return i
	} else {
		return 0
	}
}
