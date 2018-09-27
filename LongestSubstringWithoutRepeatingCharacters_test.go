package leetcode

import "testing"

var tests2 = []struct{
	input string
	expect int
} {
	{
		"abcabcbb", 3,
	},
	{
		"bbbbb", 1,
	},
	{
		"pwwkew", 3,
	},
	{
		" ", 1,
	},
	{
		"dvdf", 3,
	},
}

func TestLengthOfLongestSubstring(t *testing.T) {
	for _, te := range tests2 {
		re := lengthOfLongestSubstring1(te.input)
		if re != te.expect {
			t.Errorf("error %q. exp: %v act: %v", te, te.expect, re)
		}
	}
}
