package leetcode

func isValidSudoku(board [][]byte) bool {
	for i := 0; i < 9; i++ {
		k := make([]int, 9)
		m := make([]int, 9)
		l := make([]int, 9)
		for j := 0; j < 9; j ++ {
			if board[i][j] != '.' {
				num := board[i][j] - '0' - 1
				if k[num] == 1 {
					return false
				}
				k[num] = 1
			}
			if board[j][i] != '.' {
				num := board[j][i] - '0' - 1
				if m[num] == 1 {
					return false
				}
				m[num] = 1
			}
			x := (i / 3) * 3  + j / 3
			y := (i % 3) * 3 + j % 3
			if board[x][y] != '.' {
				num := board[x][y] - '0' - 1
				if l[num] == 1 {
					return false
				}
				l[num] = 1
			}
		}
	}

	return true
}
