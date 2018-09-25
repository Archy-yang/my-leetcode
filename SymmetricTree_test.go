package leetcode

import "testing"

var as = []struct{
	node *TreeNode
	re bool
} {
	{
		&TreeNode{
			Val:0,
		},
		true,
	},
	{
		&TreeNode{
			Val:1,
			Right:&TreeNode{
				Val:1,
			},
			Left:&TreeNode{
				Val:1,
			},
		},
		true,
	},
	{
		&TreeNode{
			Val:2,
			Right:&TreeNode{
				Val:1,
				Right:&TreeNode{
					Val:1,
				},
				Left:&TreeNode{
					Val:1,
				},
			},
			Left:&TreeNode{
				Val:1,
			},
		},
		false,
	},
}

func TestSymmetricTree(t *testing.T)  {
	for _, a := range as {
		re := isSymmetric(a.node)
		if re != a.re {
			t.Errorf("re error %v. expect %v, act %v", a.node.Val, a.re, re)
		}
	}
}
