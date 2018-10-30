package leetcode

import "testing"

var isPalindrome1Tests = []struct{
	input []int
	output bool
} {
	{
		[]int{},
		true,
	},
	{
		[]int{1},
		true,
	},
	{
		[]int{1,1},
		true,
	},
	{
		[]int{1,2,2,1},
		true,
	},
	{
		[]int{1,2},
		false,
	},
}

func TestIsPalindrome1(t *testing.T) {
	for _, test := range (isPalindrome1Tests) {
		re := isPalindrome1(tranIntListToNode(test.input))
		if re != test.output {
			t.Fatalf("failed %v %v", test, re)
		}
	}
}
