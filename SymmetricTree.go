package leetcode

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

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
		return true
	}
	rightLeft := []*TreeNode{}
	leftRight := []*TreeNode{}

	rightLeft = append(rightLeft, root.Left)
	leftRight = append(leftRight, root.Right)

	for {
		rightLeftLen := len(rightLeft)
		if rightLeftLen < 1 {
			break;
		}

		rightLeftTmp := []*TreeNode{}
		leftRightTmp := []*TreeNode{}
		for i := 0; i < rightLeftLen; i ++ {
			if rightLeft[i] == nil && leftRight[i] != nil {
				return false
			}
			if rightLeft[i] != nil && leftRight[i] == nil {
				return false
			}
			if rightLeft[i] == nil && leftRight[i] == nil {
				continue
			}
			if rightLeft[i].Val != leftRight[i].Val {
				return false
			}

			rightLeftTmp = append(rightLeftTmp, rightLeft[i].Right, rightLeft[i].Left)
			leftRightTmp = append(leftRightTmp, leftRight[i].Left, leftRight[i].Right)
		}

		if len(rightLeftTmp) != len(leftRightTmp) {
			return false
		}
		rightLeft = rightLeftTmp
		leftRight = leftRightTmp
	}

	return true
}
