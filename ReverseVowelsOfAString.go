package leetcode

import "strings"

func reverseVowels(s string) string {
	vowel := "aeiouAEIOU"
	start := 0
	end := len(s) - 1
	ss := []byte(s)
	for {
		if start >= end {
			break
		}
		if !strings.Contains(vowel, string(ss[start])) {
			start++
			continue
		}
		if !strings.Contains(vowel, string(ss[end])) {
			end--
			continue
		}
		ss[start], ss[end] = ss[end], ss[start]
		start++
		end--
	}

	return string(ss)
}
