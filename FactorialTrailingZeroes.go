package leetcode

func trailingZeroes(n int) int {
	var num int
	for num = 0; n > 0; n /= 5 {
		num += n/5
	}

	return num
}


