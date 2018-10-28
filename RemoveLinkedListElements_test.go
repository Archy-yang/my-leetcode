package leetcode

import (
	"testing"
	"reflect"
)

var removeElementsTests = []struct{
	inputList []int
	inputVal int
	output []int
} {
	{
		[]int{1,2,3,4},
		5,
		[]int{1,2,3,4},
	},
	{
		[]int{1,2,6,3,4,5,6},
		6,
		[]int{1,2,3,4,5},
	},
	{
		[]int{1},
		1,
		[]int{},
	},}

func TestRemoveElements(t *testing.T){
	for _, test := range removeElementsTests {
		re := removeElements(tranIntListToNode(test.inputList), test.inputVal)
		reList := transListNodeToIntList(re)

		if !(reflect.DeepEqual(test.output, reList)) {
			t.Fatalf("failed! %v %v", test.output, reList)
		}
	}
}