package leetcode

import "testing"
var hasPathSumTests = []struct{
	Node *TreeNode
	Sum int
	Result bool
} {
	{
		nil,
		0,
		false,
	},
	{
		&TreeNode{
			Val:0,
		},
		0,
		true,
	},
}

func TestHasPathSum(t *testing.T) {
	for _, test := range hasPathSumTests {
		re := hasPathSum(test.Node, test.Sum)
		if re != test.Result {
			t.Errorf("failed!.%q, act: %v", test, re)
		}
	}

}
