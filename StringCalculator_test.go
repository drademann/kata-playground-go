package playground

import "testing"

func TestSumEmptyStringReturns0(t *testing.T) {
	if SumOf("") != 0 {
		t.Error(`sum of "" should be 0`)
	}
}

func TestSumOfSingleNumberReturnInt(t *testing.T) {
	if SumOf("1") != 1 {
		t.Error(`sum of "1" should be 1`)
	}
}
