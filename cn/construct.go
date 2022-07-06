package cn

// Definition for a QuadTree node.
type Node struct {
	Val         bool
	IsLeaf      bool
	TopLeft     *Node
	TopRight    *Node
	BottomLeft  *Node
	BottomRight *Node
}

func construct(grid [][]int) *Node {
	l := len(grid)
	if l < 1 {
		return nil
	}
	sum := 0
	for _, item := range grid {
		for _, i := range item {
			sum += i
		}
	}
	if sum == 0 || sum == l * l {
		return &Node{
			Val: grid[0][0] == 1,
			IsLeaf: true,
		}
	}
	n := &Node{
		Val: true,
		IsLeaf: false,
	}
	tLeft := make([][]int, l/2)
	tRight := make([][]int, l/2)
	bLeft := make([][]int, l/2)
	bRight := make([][]int, l/2)
	for i := 0; i < l/2; i++ {
		tLeft[i] = grid[i][0:l/2]
		tRight[i] = grid[i][l/2:]
	}
	for i := l/2; i < l; i++ {
		bLeft[i - l/2] = grid[i][0:l/2]
		bRight[i - l/2] = grid[i][l/2:]
	}
	n.TopLeft = construct(tLeft)
	n.TopRight = construct(tRight)
	n.BottomLeft = construct(bLeft)
	n.BottomRight = construct(bRight)

	return n
}
