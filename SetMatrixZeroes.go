package leetcode

func setZeroes(matrix [][]int)  {
	r := len(matrix)
	if r < 1 {
		return
	}
	c := len(matrix[0])
	if c < 1 {
		return
	}

	var col0 bool
	for i := 0; i < r; i++ {
		if matrix[i][0] == 0 {
			col0 = true
		}
		for j := 1; j < c; j ++ {
			if matrix[i][j] == 0 {
				matrix[0][j], matrix[i][0] = 0, 0
			}
		}
	}
	for i := r-1; i >=0; i-- {
		for j := c - 1; j > 0 ; j -- {
			if matrix[0][j] == 0 || matrix[i][0] == 0  {
				matrix[i][j] = 0
			}
		}
		if col0 {
			matrix[i][0] = 0
		}
	}
}
