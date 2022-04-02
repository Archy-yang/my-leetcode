package cn

func plusOne(digits []int) []int {
	l := len(digits)
	for {
		l = l - 1
		if (l < 0) {
			break
		}
		digits[l] = digits[l] + 1
		if (digits[l] < 10) {
			break
		}
		digits[l] = digits[l] - 10
		if (l == 0) {
			digits = append(digits, 0)
			for i := 1; i < len(digits); i++ {
				digits[i-1] = digits[i]
			}
			digits[0] = 1
		}
	}
	return digits
}