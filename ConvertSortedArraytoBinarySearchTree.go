package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedArrayToBST(nums []int) *TreeNode {
	nl := len(nums)

	if nl < 1 {
		return nil
	}

	rootIndex := nl / 2
	root := &TreeNode{
		Val: nums[rootIndex],
	}
	if rootIndex > 0 {
		root.Left = sortedArrayToBST(nums[0:rootIndex])
	}
	if rootIndex + 1 < nl {
		root.Right = sortedArrayToBST(nums[rootIndex+1:nl])
	}

	return root
}