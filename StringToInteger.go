package leetcode

func myAtoi(str string) int {
	re := int(0)
	if len(str) == 0 {
		return re
	}
	f := false
	i := 0

	for ; i < len(str) ; i++ {
		if string(str[i]) != " " {
			break
		}
	}
	if i >= len(str)  || (str[i] != '-' && str[i] != '+' && (str[i] > 57 || str[i] < 48)){
		return 0
	}
	if str[i] == '-' || str[i] == '+'{
		if str[i] == '-' {
			f = true
		}
		i++
	}
	if i >= len(str) {
		return 0
	}
	if str[i] > 57 || str[i]  < 48 {
		return 0
	}
	for ; i< len(str); i++ {
		if str[i] > 57 || str[i]  < 48 {
			break;
		}
		re *= 10
		re += int(str[i] - 48)
		if re > (1 << 31) {
			break;
		}
	}

	if f {
		re *= -1
	}

	if re < -1 * (1 << 31) {
		return -1 * (1 << 31)
	}
	if re > (1 << 31 -1) {
		return 1 << 31 -1
	}
	return re
}