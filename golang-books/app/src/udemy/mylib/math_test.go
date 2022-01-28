package mylib

import "testing"

var Debug bool = false

func TestAverage(test *testing.T) {
	val := Average([]int{1, 2, 3, 4, 5, 6, 7})

	if Debug {
		test.Skip("Skip Reason")
	}

	if val != 3 {
		test.Error("Expected 3, got", test)
	}
}
