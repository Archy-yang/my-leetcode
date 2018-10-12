package leetcode

import (
	"testing"
	"reflect"
)

var getRowTest = []struct{
	Input int
	OutPut []int
} {
	{
		0,
		[]int{1},
	},
	{
		1,
		[]int{1,1},
	},
	{
		2,
		[]int{1,2,1},
	},
	{
		4,
		[]int{1,4,6,4,1},
	},
}

func TestGetRow(t *testing.T) {
	for _, test := range getRowTest {
		re := getRow(test.Input)

		if !reflect.DeepEqual(re, test.OutPut) {
			t.Errorf("failed. %q, act:%v", test, re)
		}
	}
}

