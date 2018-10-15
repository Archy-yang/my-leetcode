package leetcode

import "testing"

var singleNumber2Tests = []struct{
	input []int
	output int
} {
	//{
	//	[]int{2,2,1,2},
	//	1,
	//},
	//{
	//	[]int{0,1,0,1,0,1,99},
	//	99,
	//},
	{
		[]int{2,3,2,3,2,3,1},
		4,
	},
	//{
	//	[]int{-2,-2,1,1,-3,1,-3,-3,-4,-2},
	//	-4,
	//},
}

func TestSingleNumber2(t *testing.T) {
	for _, test := range singleNumber2Tests {
		re := singleNumber2(test.input)

		if re != test.output {
			t.Fatalf("faild!, %q, %v", test, re)
		}
	}
}