package leetcode

func generate(numRows int) [][]int {
	result := make([][]int, numRows)
	if numRows == 0 {
		return result
	}
	result[0] = []int{1}
	if numRows == 1 {
		return result
	}
	result[1] = []int{1,1}
	if numRows == 2 {
		return result
	}
	for i := 2; i < numRows; i++ {
		t := make([]int, i+1)
		t[0] = 1
		t[i] = 1
		for  j := 1; j < i; j ++ {
			t[j] =  result[i-1][j-1] + result[i-1][j]
		}
		result[i] = t
	}

	return result
}
