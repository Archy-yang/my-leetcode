package leetcode

func reverseString(s string) string {
	r := ""
	for i := 0; i < len(s); i ++ {
		r = s[i:i+1] + r
	}

	return r
}

func reverseString2(s string) string {
	if len(s) <= 1 {
		return s
	}
	m := len(s) / 2

	l := reverseString2(s[0:m])
	r := reverseString2(s[m:len(s)])

	return r+l
}
