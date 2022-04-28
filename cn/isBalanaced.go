package cn

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {
	b, _ := dsf(root)
	return b
}

func dsf(root *TreeNode) (bool, int) {
	if root == nil {
		return true, 0
	}
	b, left := dsf(root.Left)
	if !b {
		return false, 0
	}
	b, right := dsf(root.Right)
	if !b {
		return false, 0
	}
	if left - right > 1 || right - left > 1 {
		return false, 0
	}
	left++
	right++
	if left > right {
		return true, left
	}
	return true, right
}
