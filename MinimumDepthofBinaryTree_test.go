package leetcode

import "testing"

var minDepthTests = []struct{
	Node *TreeNode
	Result int
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
			Left: &TreeNode{
				Val:2,
			},
		},
		2,
	},
	{
		&TreeNode{
			Val:1,
			Right: &TreeNode{
				Val:2,
			},
		},
		2,
	},
	{
		&TreeNode{
			Val:1,
			Left: &TreeNode{
				Val:2,
			},
			Right: &TreeNode{
				Val:2,
				Left:&TreeNode{
					Val:3,
				},
				Right:&TreeNode{
					Val:3,
				},
			},
		},
		2,
	},
}

func TestMinDepth(t *testing.T) {
	for _, test := range minDepthTests {
		re := minDepth(test.Node)

		if re != test.Result {
			t.Errorf("fail. %v, act:%v", test, re)
		}
	}
}

func TestMinDepth1(t *testing.T) {
	for _, test := range minDepthTests {
		re := minDepth1(test.Node)

		if re != test.Result {
			t.Errorf("fail. %v, act:%v", test, re)
		}
	}
}
