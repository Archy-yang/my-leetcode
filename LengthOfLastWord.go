package leetcode

import "strings"

func lengthOfLastWord(s string) int {
	wlen := 0
	s = strings.Trim(s, " ")

	for i := 0; i < len(s); i++ {
		if s[i] != ' ' {
			wlen++
		} else {
			wlen = 0
		}
	}

	return wlen
}
