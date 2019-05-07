package leetcode

func spiralOrder(matrix [][]int) []int {
	m := len(matrix)
	if m < 1 {
		return []int{}
	}
	n := len(matrix[0])
	if n < 1 {
		return []int{}
	}

	rowStart, rowMax := 0, m-1
	colStart, colMax := 0, n-1

	re := []int{}
	for rowMax > rowStart && colMax > colStart {
		for j := colStart; j < colMax; j++ {
			re = append(re, matrix[rowStart][j])
		}
		for i := rowStart;i < rowMax; i++ {
			re = append(re, matrix[i][colMax])
		}
		for j := colMax;j > colStart; j-- {
			re = append(re, matrix[rowMax][j])
		}
		for i := rowMax;i > rowStart; i-- {
			re = append(re, matrix[i][colStart])
		}
		rowMax--
		rowStart++
		colMax--
		colStart++
	}

	if rowMax == rowStart {
		for j := colStart; j <= colMax; j ++ {
			re = append(re, matrix[rowMax][j])
		}
	} else if colMax == colStart {
		for i := rowStart;i <= rowMax; i ++ {
			re = append(re, matrix[i][colMax])
		}
	}

	return re
}