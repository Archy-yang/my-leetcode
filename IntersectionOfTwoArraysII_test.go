package leetcode

import (
	"testing"
	"reflect"
)

var intersectTests = []struct{
	in1 []int
	in2 []int
	out []int
}{
	{
		[]int{1,2,3},
		[]int{1,2,3},
		[]int{1,2,3},
	},
	{
		[]int{1,2,2,1},
		[]int{2,2},
		[]int{2,2},
	},
	{
		[]int{4,9,5},
		[]int{9,4,9,8,4},
		[]int{9,4},
	},
}
func TestIntersect(t *testing.T)  {
	for _, test := range intersectTests {
		re := intersect(test.in1, test.in2)
		if !reflect.DeepEqual(re, test.out) {
			t.Fatalf("failed! exp: %v, act: %v", test.out, re)
		}
	}
}