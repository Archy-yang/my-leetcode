package leetcode

func addDigits(num int) int {
	for num > 9 {
		n := 0
		t := num
		for t != 0 {
			n += t%10
			t /= 10
		}

		num = n
	}

	return num
}
