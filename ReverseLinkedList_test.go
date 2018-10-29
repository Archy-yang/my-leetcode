package leetcode

import (
	"testing"
	"reflect"
)

var reverseListTests = []struct{
	intput, output []int
} {
	{
		[]int{1,2,3,4,5},
		[]int{5,4,3,2,1},
	},
	{
		[]int{},
		[]int{},
	},
}

func TestReverseList(t *testing.T) {
	for _, test := range reverseListTests {
		re := reverseList(tranIntListToNode(test.intput))
		reList := transListNodeToIntList(re)

		if !reflect.DeepEqual(reList, test.output) {
			t.Fatalf("failed !! %v, %v", test, reList)
		}
	}
}
func TestReverseList1(t *testing.T) {
	for _, test := range reverseListTests {
		re := reverseList1(tranIntListToNode(test.intput))
		reList := transListNodeToIntList(re)

		if !reflect.DeepEqual(reList, test.output) {
			t.Fatalf("failed !! %v, %v", test, reList)
		}
	}
}
