package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}

	return help(root, 0, sum)
}

func help(root *TreeNode, add int, sum int) bool {
	add += root.Val
	if root.Left == nil && root.Right == nil && add == sum {
		return true
	}
	if root.Left != nil {
		lRe := help(root.Left, add, sum)
		if lRe {
			return true
		}
	}
	if root.Right != nil {
		rRe := help(root.Right, add, sum)
		if rRe {
			return true
		}
	}

	return false
}