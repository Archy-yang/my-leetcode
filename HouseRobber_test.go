package leetcode

import "testing"

var robTests = []struct {
	input []int
	result int
} {
	{
		[]int{1,2,3,1},
		4,
	},
	{
		[]int{2,7,9,3,1},
		12,
	},
	{
		[]int{2,1,1,2},
		4,
	},
}

func TestRob(t *testing.T) {
	for _, test := range robTests {
		re := rob(test.input)

		if re != test.result {
			t.Fatalf("failed!! %v %v", test, re)
		}
	}
}
