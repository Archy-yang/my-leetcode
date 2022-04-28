package cn

// Definition for a binary tree node.
 type TreeNode struct {
     Val int
     Left *TreeNode
     Right *TreeNode
 }

func inorderTraversal(root *TreeNode) []int {
    n := []int{}
    if root == nil {
        return n
    }
    return depth(root, n)
}

func depth(t *TreeNode, n []int) []int {
    if t.Left != nil {
        n = append(n, depth(t.Left, []int{})...)
    }
    n = append(n, t.Val)
    if t.Right != nil {
        n = append(n, depth(t.Right, []int{})...)
    }
    return n
}
