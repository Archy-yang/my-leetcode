package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrderBottom(root *TreeNode) [][]int {
	var orders [][]int

	if root != nil {
		var nodes []*TreeNode
		nodes = append(nodes, root)
		for {
			tmp := nodes
			if len(tmp) == 0 {
				break
			}
			nodes = []*TreeNode{}
			var order []int

			for _, node := range tmp {
				order = append(order, node.Val)
				if node.Left != nil {
					nodes = append(nodes, node.Left)
				}
				if node.Right != nil {
					nodes = append(nodes, node.Right)
				}
			}

			orders = append(orders, order)
		}

		l := len(orders)
		c := l/2
		for i:=0; i < c; i++ {
			orders[i], orders[l - i - 1] = orders[l - i - 1], orders[i]
		}
	}

	return orders
}
