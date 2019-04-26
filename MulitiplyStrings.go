package leetcode

// @TODO 可以调整成申请[]btye{}
func multiply(num1 string, num2 string) string {
	l1 := len(num1)
	l2 := len(num2)
	if l1 < 1 || l2 < 1 {
		return ""
	}
	if num1[0] == '0' || num2[0] == '0' {
		return "0"
	}
	r := ""
	s := 0
	for i := 1; i < (l1 + l2); i++ {
		t := 0
		p1 := 1
		if i > l2 {
			p1 = i + 1 - l2
		}
		for ; p1 <= l1 && p1 <= i; p1++ {
			p2 := i + 1 - p1
			n1 := int(num1[l1 - p1] - '0')
			n2 := int(num2[l2 - p2] - '0')

			t += n1 * n2
		}
		 t += s

		 s = t / 10
		 r = string((t % 10) + '0') + r
	}
	for s > 0 {
		r = string( s % 10 + '0') + r
		s /= 10
	}

	return r
}
