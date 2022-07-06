package cn

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func largestValues(root *TreeNode) []int {
	if root == nil {
		return []int {}
	}
	r := []int{}

	max := root.Val

	q := []*TreeNode{}
	p, p1 := 0, 0
	l := 1

	q = append(q, root)
	for ; p < l; p++ {
		if q[p].Left != nil {
			q = append(q, q[p].Left)
			l++
		}
		if q[p].Right != nil {
			q = append(q, q[p].Right)
			l++
		}
		if max < q[p].Val {
			max = q[p].Val
		}
		if p == p1 {
			p1 = l - 1
			r = append(r, max)
			max = q[p1].Val
		}
	}

	return r
}
