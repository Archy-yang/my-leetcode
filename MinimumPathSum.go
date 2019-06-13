package leetcode

func minPathSum(grid [][]int) int {
	row := len(grid)
	if row == 0 {
		return 0
	}
	col := len(grid[0])
	if col == 0 {
		return 0
	}

	tmp := make([]int, row)

	for i := 0 ; i < col; i++ {
		for j := 0; j < row; j++ {
			if j == 0 || (i > 0 && tmp[j-1] >= tmp[j]) {
				tmp[j] += grid[j][i]
				continue
			}
			if i == 0 || tmp[j-1] < tmp[j] {
				tmp[j] = tmp[j-1] + grid[j][i]
				continue
			}
		}
	}
	return tmp[row - 1]
}
