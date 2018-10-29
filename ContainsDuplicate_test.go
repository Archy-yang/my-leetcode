package leetcode

import "testing"

var containsDuplicateTests = []struct{
	input []int
	output bool
} {
	{
		[]int{1,2,3,1},
		true,
	},
	{
		[]int{1,2,3},
		false,
	},
}

func TestContainsDuplicate(t *testing.T) {
	for _, test := range containsDuplicateTests {
		re := containsDuplicate(test.input)

		if re != test.output{
			t.Fatalf("failed!!, %v, %v", test, re)
		}
	}
}
