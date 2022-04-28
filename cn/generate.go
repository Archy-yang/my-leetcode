package cn

func generate(numRows int) [][]int {
	if numRows == 1 {
		return [][]int{{1}}
	}
	n := [][]int{{1}, {1, 1}}
	for i := 2; i < numRows; i++ {
		t := []int{1}
		for j := 1; j < i; j++ {
			t = append(t, n[i-1][j-1]+n[i-1][j])
		}
		t = append(t, 1)
		n = append(n, t)
	}
	return n
}

func getRow(rowIndex int) []int {
	if rowIndex == 0 {
		return []int{1}
	}
	n := make([]int, rowIndex+1)
	n[0], n[1] = 1, 1
	for i := 2; i <= rowIndex; i++ {
		n[i] = 1
		for j := i - 1; j > 0; j-- {
			n[j] = n[j-1] + n[j]
		}
	}
	return n
}
