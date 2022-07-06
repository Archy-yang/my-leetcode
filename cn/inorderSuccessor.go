package cn

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {
	l := ergodic(root)
	ll := len(l)
	i:=0
	for ; i < ll; i++ {
		if l[i] == p {
			break
		}
	}
	i++
	if i >= ll {
		return nil
	}
	return l[i]
}


func inorderSuccessor2(root *TreeNode, p *TreeNode) *TreeNode {
	var prev *TreeNode
	for ;root.Val != p.Val; {
		if root.Val < p.Val {
			root = root.Right
		} else {
			prev = root
			root = root.Left
		}
	}
	if root.Right != nil {
		t := root.Right
		for ; t.Left != nil; {
			t = t.Left
		}
		return t
	}

	return prev
}

func ergodic(root *TreeNode) []*TreeNode {
	l := []*TreeNode{}
	if root == nil {
		return l
	}
	if root.Left != nil {
		l = append(l, ergodic(root.Left)...)
	}
	l = append(l, root)
	if root.Right != nil {
		l = append(l, ergodic(root.Right)...)
	}
	return l
}
