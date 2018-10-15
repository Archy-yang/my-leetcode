package leetcode

import "strings"

func isPalindrome(s string) bool {
	s = strings.ToUpper(s)
	i := 0
	j := len(s)-1
	for {
		if j < i {
			break
		}
		if s[i] < 48 || (s[i] > 57 && s[i] < 65) || s[i] > 90 {
			i++
			continue
		}
		if s[j] < 48 || (s[j] > 57 && s[j] < 65) || s[j] > 90 {
			j--
			continue
		}
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}

	return true
}
