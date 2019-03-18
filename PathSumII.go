package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, sum int) [][]int {
	var dfs func(node *TreeNode, sum, target int, list []int)
	re := [][]int{}

	dfs = func(root *TreeNode, sum, target int, list []int) {
		if root == nil {
			return
		}
		sum += root.Val
		list = append(list, root.Val)
		if root.Left == nil && root.Right == nil && sum == target {
			re = append(re, list)
			return
		}
		leftList := make([]int, len(list))
		copy(leftList, list)
		rightList := list
		dfs(root.Left, sum, target, leftList)
		dfs(root.Right, sum, target, rightList)
	}

	dfs(root, 0, sum, []int{})
	return re
}
