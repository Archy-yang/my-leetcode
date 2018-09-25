package leetcode

func plusOne(digits []int) []int {
	length := len(digits)
	LabelOut:
	for i := length - 1; i>=0; i-- {
		if digits[i] + 1 < 10 {
			digits[i] += 1
			break LabelOut
		} else {
			digits[i] = 0
			if i == 0 {
				tmp := []int{1}
				digits = append(tmp, digits...)
			}
		}
	}

	return digits
}
