package leetcode

import (
	"testing"
	"reflect"
)

var twoSumTests = []struct{
	input, output []int
	target int
} {
	{
		[]int{2,7,11,15},
		[]int{1,2},
		9,
	},
	{
		[]int{2,7,11,15},
		[]int{1,3},
		13,
	},
}
func TestTwoSum(t *testing.T) {
	for _, test := range twoSumTests {
		re := twoSum(test.input, test.target)

		if !reflect.DeepEqual(re, test.output) {
			t.Fatalf("failed! %v %v", test, re)
		}
	}
}
