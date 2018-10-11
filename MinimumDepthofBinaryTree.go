package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftD := minDepth(root.Left)
	rightD := minDepth(root.Right)

	if rightD == 0 || (leftD != 0 && leftD < rightD) {
		return leftD + 1
	}
	return rightD + 1
}

func minDepth1(root *TreeNode) int {
	depth := 0
	var nodes []*TreeNode
	if root != nil {
		nodes = append(nodes, root)
	}
	for {
		if len(nodes) < 1 {
			break
		}
		depth++
		tmpNodes := []*TreeNode{}
		for _, node := range nodes {
			if node.Right == nil && node.Left == nil {
				return depth
			}
			if node.Left != nil {
				tmpNodes = append(tmpNodes, node.Left)
			}
			if node.Right != nil {
				tmpNodes = append(tmpNodes, node.Right)
			}
		}
		nodes = tmpNodes
	}

	return depth
}