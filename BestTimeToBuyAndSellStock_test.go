package leetcode

import (
	"testing"
)

var maxProfitTests = []struct{
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
		[]int{7,1,5,3,6,4},
		5,
	},
}

func TestMaxProfix(t *testing.T) {
	for _, test := range maxProfitTests {
		re := maxProfit(test.Input)

		if re != test.Output {
			t.Errorf("failed ! %q act %v", test, re)
		}
	}
}
