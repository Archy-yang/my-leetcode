package cn

func findDiagonalOrder(mat [][]int) []int {
	m, n := len(mat), len(mat[0])
	t := []int{}
	for i := 0; i < m + n -1; i++ {
		if i % 2 == 0 {
			y := maxFunc(0, i + 1 - m)
			x := minFunc(i, m-1)
			for ;y <= n-1 && x >=0; {
				t = append(t, mat[x][y])
				x -= 1
				y += 1
			}
		} else {
			y := minFunc(i, n - 1)
			x := maxFunc(0, i + 1 - n)
			for ; y >= 0 && x <= m-1; {
				t = append(t, mat[x][y])
				x += 1
				y -= 1
			}
		}
	}
	return t
}

func maxFunc(a, b int ) int {
	if a > b {
		return a
	}
	return b
}

func minFunc(a, b int ) int {
	if a < b {
		return a
	}
	return b
}
