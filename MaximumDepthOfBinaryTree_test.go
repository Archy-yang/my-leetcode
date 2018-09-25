package leetcode

import "testing"

var trs = []struct{
	node *TreeNode
	dep int
} {
	{
		nil,
		0,
	},
	{
		&TreeNode{
			Val:1,
		},
		1,
	},
	{
		&TreeNode{
			Val:1,
			Left:&TreeNode{
				Val:1,
			},
		},
		2,
	},
	{
		&TreeNode{
			Val:1,
			Left:&TreeNode{
				Val:1,
			},
			Right:&TreeNode{
				Val:1,
				Left:&TreeNode{
					Val:1,
				},
			},
		},
		3,
	},
}

func TestMaximumDepthOfBinaryTerr(t *testing.T)  {
	for _, tr := range trs {
		re := maxDepth(tr.node)
		if re != tr.dep {
			t.Errorf("re error %v. expect %v, act %v", tr.node.Val, tr.dep, re)
		}
	}
}
