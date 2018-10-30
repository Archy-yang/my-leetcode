package leetcode

import "testing"

var isPowerOfTwoTests = []struct{
	input int
	output bool
} {
	{
		0,false,
	},
	{
		1,true,
	},
	{
		2,true,
	},
	{
		16,true,
	},
	{
		218,false,
	},
}

func TestIsPowerOfTwo(t *testing.T) {
	for _, test := range isPowerOfTwoTests {
		re := isPowerOfTwo(test.input)

		if re != test.output {
			t.Fatalf("failed!! %v %v", test, re)
		}
	}
}
