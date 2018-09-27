package leetcode

import "strings"

func lengthOfLongestSubstring(s string) int {
	var l int
	sub := map[uint8]int{}
	sl := len(s)

	for i := 0; i < sl ; i++ {
		if oi, ok := sub[s[i]]; ok {
			if l < len(sub) {
				l = len(sub)
			}
			i = oi+1
			sub = map[uint8]int{}

		}
		sub[s[i]] = i
	}
	if l < len(sub) {
		l = len(sub)
	}

	return l
}

func lengthOfLongestSubstring1(s string) int {
	var l int
	sub := ""

	for i, _ := range s {
		c := s[i:i+1]
		if f := strings.Index(sub, c); f != -1 {
			if l < len(sub) {
				l = len(sub)
			}

			sub = sub[f+1:]
		}
		sub += c
	}
	if l < len(sub) {
		l = len(sub)
	}

	return l
}
