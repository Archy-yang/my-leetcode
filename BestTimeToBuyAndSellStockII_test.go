package leetcode

import (
	"testing"
)

var maxProfitIITests = []struct{
	Input []int
	Output int
} {
	{
		[]int{},
		0,
	},
	{
		[]int{7,6,4,3,1},
		0,
	},
	{
		[]int{1,2,3,4,5},
		4,
	},
	{
		[]int{7,1,5,3,6,4},
		7,
	},
}

func TestMaxProfixII(t *testing.T) {
	for _, test := range maxProfitIITests {
		re := maxProfitII(test.Input)

		if re != test.Output {
			t.Errorf("failed ! %q exp: %v, act %v", test, test.Output, re)
		}
	}
}