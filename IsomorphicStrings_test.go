package leetcode

import "testing"

var isIsomorphicTests = []struct{
	inputS, inputT string
	output bool
} {
	{
		"a","b", true,
	},
	{
		"aa","bb", true,
	},
	{
		"ab","bb", false,
	},
	{
		"","", true,
	},
	{
		"egg","add", true,
	},
	{
		"foo","bar", false,
	},
	{
		"paper","title", true,
	},
}

func TestIsIsomorphic (t *testing.T) {
	for _, test := range isIsomorphicTests {
		re := isIsomorphic(test.inputS, test.inputT)
		if re != test.output {
			t.Fatalf("failed %v %v", test, re)
		}
	}
}
