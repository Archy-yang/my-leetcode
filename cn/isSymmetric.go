package cn

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return false
	}
	return symmetric(root.Left, root.Right)
}

func symmetric(t1, t2 *TreeNode) bool {
	if t1 == nil && t2 == nil {
		return true
	}
	if t1 == nil || t2 == nil {
		return false
	}
	if t1.Val != t2.Val {
		return false
	}
	if !symmetric(t1.Left, t2.Right) {
		return false
	}
	if !symmetric(t1.Right, t2.Left) {
		return false
	}
	return true
}
