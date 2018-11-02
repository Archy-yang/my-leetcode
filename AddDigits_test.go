package leetcode

import "testing"

var addDigitsTests = []struct{
	input, output int
} {
	{
		38, 2,
	},
	{
		10, 1,
	},
}

func TestAddDigits(t *testing.T) {
	for _, test := range addDigitsTests {
		re := addDigits(test.input)

		if re != test.output {
			t.Fatalf("failed !! %v %v", test, re)
		}
	}
}
