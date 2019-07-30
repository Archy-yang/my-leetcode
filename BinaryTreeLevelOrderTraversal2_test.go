package leetcode

import (
	"testing"
)

var tr = []struct{
	node *TreeNode
	order [][]int
} {
	{
		nil,
		[][]int{},
	},
	{
		&TreeNode{
			Val:1,
		},
		[][]int{
			{1},
		},
	},
	{
		&TreeNode{
			Val:1,
			Left:&TreeNode{
				Val:1,
			},
			Right:&TreeNode{
				Val:1,
			},
		},
		[][]int{
			{1},
			{1,1},
		},
	},
}

func Test_levelOrderBottom(t *testing.T) {
	for _, te := range tr {
		re := levelOrderBottom(te.node)

		if len(re) != len(te.order) {
			t.Errorf("error: %v", te)
		}
	}
}




