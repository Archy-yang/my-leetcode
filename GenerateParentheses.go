package leetcode

var re []string
func generateParenthesis(n int) []string {
	re = []string{}
	gen("", n, n)

	return re
}

func gen(s string, left, right int) {
	if left == 0 && right ==0 {
		re = append(re, s)
		return
	}
	if left >= right || left > 0 {
		gen(s + "(", left - 1, right)
	}
	if right > left && right > 0 {
		gen(s + ")", left, right - 1)
	}
}
