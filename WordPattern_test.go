package leetcode

import "testing"

var wordPatternTests = []struct{
	inputP, inputS string
	output bool
} {
	{
		"abba", "dog cat cat dog", true,
	},
	{
		"abba", "dog cat cat fish", false,
	},
	{
		"aaaa", "dog dog dog dog", true,
	},
	{
		"ab", "dog dog", false,
	},
	{
		"aa", "dog dog", true,
	},
}

func TestWordPattern(t *testing.T) {
	for _, test := range wordPatternTests {
		re := wordPattern(test.inputP, test.inputS)

		if re != test.output {
			t.Fatalf("failed!! %v %v", test, re)
		}
	}
}
