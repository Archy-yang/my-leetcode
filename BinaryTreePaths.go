package leetcode

import "strconv"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return []string{}
	}
	left := binaryTreePaths(root.Left)
	right := binaryTreePaths(root.Right)

	r := strconv.Itoa(root.Val)
	l := append(left, right...)

	if len(l) < 1 {
		return []string{r}
	}
	for i, p := range l {
		l[i] = r + "->" + p
	}
	return l
}
