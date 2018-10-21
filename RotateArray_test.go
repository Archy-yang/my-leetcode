package leetcode

import (
	"testing"
	"reflect"
)

var rotateTests = []struct{
	intput, output []int
	k int
} {
	{
		intput: []int{1,2,3,4,5,6,7},
		k: 3,
		output: []int{5,6,7,1,2,3,4},
	},
}

func TestRotateArray(t *testing.T) {
	for _, test := range rotateTests {
		rotate(test.intput, test.k)
		if !reflect.DeepEqual(test.output, test.intput) {
			t.Fatalf("failed, %v",test)
		}
	}
}
