package leetcode

import (
	"testing"
	"reflect"
)

var pascalsTriangleTest = []struct{
	Input int
	OutPut [][]int
} {
	{
		0,
		[][]int{},
	},
	{
		1,
		[][]int{
			{1},
		},
	},
	{
		2,
		[][]int{
			{1}, {1,1},
		},
	},
	{
		3,
		[][]int{
			{1}, {1,1},{1,2,1},
		},
	},
	{
		5,
		[][]int{
			{1}, {1,1},{1,2,1},{1,3,3,1}, {1,4,6,4,1},
		},
	},
}

func TestPascalsTriangle(t *testing.T) {
	for _, test := range pascalsTriangleTest {
		re := generate(test.Input)

		if !reflect.DeepEqual(re, test.OutPut) {
			t.Errorf("failed. %q, act:%v", test, re)
		}
	}
}
