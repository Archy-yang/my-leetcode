package cn

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedArrayToBST(nums []int) *TreeNode {
	l := len(nums)
	if l == 1 {
		return &TreeNode{Val: nums[0]}
	}

	m := l / 2
	t := &TreeNode{Val: nums[m]}
	t.Right = sortedArrayToBST(nums[m+1:])
	t.Left = sortedArrayToBST(nums[0:m])

	return t
}
