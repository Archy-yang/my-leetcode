package leetcode

func isPerfectSquare(num int) bool {
	if num == 1 {
		return true
	}
	for i := 1; i <= num / 2 ; i ++ {
		if i * i == num {
			return true
		}
	}
	return false
}

func isPerfectSquare2(num int) bool {
	if num == 1 {
		return true
	}
	a, b := 1, num
	for  a < b {
		m := (a + b) / 2
		m2 := m * m
		if m2 == num {
			return true
		} else if (m2 > num) {
			b = m
		} else {
			if a == m {
				return false
			}
			a = m
		}
	}
	return false
}
