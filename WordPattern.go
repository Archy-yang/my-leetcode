package leetcode

import "strings"

func wordPattern(pattern string, str string) bool {
	strs := strings.Split(str, " ")

	if len(strs) != len(pattern) {
		return false
	}
	pm := make([]int, 26)
	sm := make(map[string]int, len(strs))

	for i, b := range pattern {
		pi := int(b) - int('a')
		if si, ok := sm[strs[i]]; (!ok && si > 0) || si != pm[pi] {
			return false
		}
		pm[pi] = i + 1
		sm[strs[i]] = i + 1
	}

	return true
}
