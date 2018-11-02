package leetcode

import "testing"

var isAnagramTests = []struct{
	inputS, inputT string
	output bool
} {
	{
		"anagram", "nagaram", true,
	},
	{
		"rat", "car", false,
	},
}

func TestIsAnagram(t *testing.T) {
	for _, test := range isAnagramTests {
		re := isAnagram(test.inputS, test.inputT)

		if re != test.output {
			t.Fatalf("failed! %v %v", test, re)
		}
	}
}

