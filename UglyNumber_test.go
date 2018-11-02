package leetcode

import "testing"

var isUglyTests = []struct{
	input int
	output bool
} {
	{
		6, true,
	},
	{
		14, false,
	},
	{
		1, true,
	},
	{
		2, true,
	},
}

func TestIsUgly(t *testing.T) {
	for _, test := range isUglyTests {
		re := isUgly(test.input)

		if re != test.output {
			t.Fatalf("failed!! %v %v", test, re)
		}
	}
}
