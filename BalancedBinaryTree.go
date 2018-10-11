package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {
	_, re := search(root)

	return re
}

func search(root *TreeNode) (int, bool) {
	if root == nil {
		return 0, true
	}
	leftL := 0
	rightL := 0
	leftL, lb := search(root.Left)
	if !lb {
		return 0, lb
	}
	rightL, rb := search(root.Right)
	if !rb {
		return 0, rb
	}

	d := leftL - rightL
	if d > 1 || d < -1 {
		return 0, false
	}

	if d > 0 {
		return leftL + 1, true
	}
	return rightL + 1, true
}