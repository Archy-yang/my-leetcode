package leetcode

func longestPalindrome(s string) string {
	max := ""
	for i := 0; i < len(s); i ++ {
		for j := i; j < len(s); j++ {
			f := true
			start, end := i, j
			for start <= end {
				if s[start] != s[end] {
					f = false
					break;
				}
				start++
				end--
			}

			if f {
				if len(max) <= (j - i) {
					max = string(s[i:j+1])
				}
			}
		}
	}

	return max
}
