package leetcode

import (
	"testing"
	"reflect"
)

var moveZeroesTests = []struct{
	input []int
	output []int
} {
	{
		[]int{0,1,0,3,12},
		[]int{1,3,12,0,0},
	},
	{
		[]int{0,0,1},
		[]int{1,0,0},
	},
}

func TestMoveZeroes(t *testing.T) {
	for _, test := range moveZeroesTests {
		moveZeroes(test.input)

		if !reflect.DeepEqual(test.input, test.output) {
			t.Fatalf("failed %v", test)
		}
	}
}
