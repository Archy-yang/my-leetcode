package leetcode

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	rowNum := len(obstacleGrid)
	if rowNum < 1 {
		return 0
	}
	columnNum := len(obstacleGrid[0])
	if columnNum < 1 {
		return 0
	}
	tmp := make([]int, rowNum)
	for j:= 0; j < columnNum; j ++ {
		for i:= 0; i < rowNum; i++ {
			if obstacleGrid[i][j] != 0 {
				tmp[i] = 0
			} else {
				if j == 0 && i == 0{
					tmp[i] = 1
				} else {
					if i > 0 {
						tmp[i] += tmp[i-1]
					}
				}
			}
		}
	}

	return tmp[rowNum-1]
}
