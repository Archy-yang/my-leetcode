package leetcode

func longestValidParentheses(s string) int {
	var max = 0
	for i := 0; i < len(s); {
		if s[i] == ')' {
			i++
			continue
		}
		t := 1
		n := 1
		m := 1
		for j := i +1; j < len(s); j++ {
			n++
			if s[j] == ')' {
				t--
			} else {
				t++
			}
			if t < 0 {
				break;
			}
			if t == 0 && max < n {
				max = n
				m = n
			}
		}
		i += m
	}

	return max
}

