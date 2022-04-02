package cn

func lengthOfLastWord(s string) int {
	l := len(s) - 1
	flag := false
	c := 0
	for ; l >= 0; l-- {
		if (flag && s[l] == ' ') {
			break
		}
		if s[l] != ' ' {
			flag = true
			c++
			continue
		}
	}

	return c
}
