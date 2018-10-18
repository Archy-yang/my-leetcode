package leetcode

func convertToTitle(n int) string {
	t := ""
	for {
		c := n % 26
		n = n / 26
		if c == 0 {
			c = 26
			n--
		}
		c += 64
		cStr := string([]byte{byte(c)})
		t = cStr + t

		if n == 0 {
			break
		}
	}
	return t
}
