package leetcode

import "testing"

var singleNumberTests = []struct{
	input []int
	output int
} {
	{
		[]int{2,2,1},
		1,
	},
}

func TestSingleNumber(t *testing.T) {
	for _, test := range singleNumberTests {
		re := singleNumber(test.input)

		if re != test.output {
			t.Errorf("faild!, %q, %v", test, re)
		}
	}
}
