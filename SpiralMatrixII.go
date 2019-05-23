package leetcode

func generateMatrix(n int) [][]int {
	if n == 0 {
		return [][]int{}
	}
	if n == 1 {
		return [][]int{{1}}
	}
	all := make([][]int, n)
	for i := range all {
		all[i] = make([]int, n)
	}

	k := 1
	min, max := 0, n - 1
	for min < max{
		for j := min ; j < max; j++ {
			all[min][j] = k
			k++
		}
		for i := min; i < max; i++ {
			all[i][max] = k
			k++
		}
		for j := max; j > min; j-- {
			all[max][j] = k
			k++
		}
		for i := max; i > min; i-- {
			all[i][min] = k
			k++
		}
		min++
		max--
	}

	if min == max {
		all[min][max] = k
	}

	return all
}
