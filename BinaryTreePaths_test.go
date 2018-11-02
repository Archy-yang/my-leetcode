package leetcode

import (
	"testing"
	"reflect"
)

var binaryTreePathsTests = []struct{
	input *TreeNode
	output []string
} {
	{
		&TreeNode{
			Val:1,
			Left:nil,
			Right:nil,
		},
		[]string{"1"},
	},
	{
		&TreeNode{
			Val:1,
			Left:&TreeNode{
				Val:2,
			},
			Right:&TreeNode{
				Val:3,
			},
		},
		[]string{"1->2", "1->3"},
	},
}

func TestBinaryTreePaths(t *testing.T) {
	for _, test := range binaryTreePathsTests{
		re := binaryTreePaths(test.input)

		if !reflect.DeepEqual(re, test.output) {
			t.Fatalf("failed! %v %v", test, re)
		}
	}
}
