package leetcode

func isUgly(num int) bool {
	if num < 1 {
		return false
	}
	for num != 1 && num != 2 && num != 3 && num != 5 {
		if num % 2 == 0 {
			num /= 2
			continue
		}
		if num % 3 == 0 {
			num /= 3
			continue
		}
		if num % 5 == 0 {
			num /= 5
			continue
		}
		return false
	}
	return true
}

