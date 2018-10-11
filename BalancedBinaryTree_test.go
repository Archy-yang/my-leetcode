package leetcode

import "testing"

var testData = []struct{
	Node *TreeNode
	Result bool
}{
	{
		nil,
		true,
	},
	{
		&TreeNode{
			Val:1,
		},
		true,
	},
	{
		&TreeNode{
			Val:1,
			Left:&TreeNode{
				Val: 2,
				Left:&TreeNode{
					Val:3,
				},
			},
		},
		false,
	},
}

func TestIsBalanced (t *testing.T) {
	for _, td := range testData {
		re := isBalanced(td.Node)

		if re != td.Result {
			t.Errorf("error: %q, act: $v", td, re)
		}
	}
}
