package leetcode

import (
	"testing"
	"reflect"
)

var intersectionTests = []struct{
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
		[]int{2},
	},
}
func TestIntersection(t *testing.T)  {
	for _, test := range intersectionTests {
		re := intersection(test.in1, test.in2)
		if !reflect.DeepEqual(re, test.out) {
			t.Fatalf("failed! exp: %v, act: %v", test.out, re)
		}
	}
}
