package leetcode

func rotateImage(matrix [][]int)  {
	if len(matrix) < 2 {
		return
	}
	ll := len(matrix)

	for i := 0; i < ll; i ++ {
		for j := 0; j < i; j ++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	for i := 0; i < ll; i++ {
		for j := 0; j < ll / 2; j++ {
			matrix[i][j], matrix[i][ll-1-j] = matrix[i][ll-1-j], matrix[i][j]
		}
	}
}
