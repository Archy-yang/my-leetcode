package leetcode

import "testing"

var isHappyTests = []struct{
	input int
	output bool
} {
	{
		19, true,
	},
	{
		82, true,
	},
}

func TestIsHappy(t *testing.T) {
	for _, test := range isHappyTests{
		re := isHappy(test.input)

		if re != test.output {
			t.Fatalf("failed! %v, %v", test, re)
		}
	}
}
