package leetcode

func getRow(rowIndex int) []int {
	row := make([]int, rowIndex+1)
	row[0] = 1
	if rowIndex == 0 {
		return row
	}
	row[1] = 1
	if rowIndex == 1 {
		return row
	}
	for i := 2; i < rowIndex+1; i++ {
		tmpRow := make([]int, i+1)
		tmpRow[0] = 1
		tmpRow[i] = 1
		for j := 1; j < i; j ++ {
			tmpRow[j] = row[j-1] + row [j]
		}

		row = tmpRow
	}

	return row
}