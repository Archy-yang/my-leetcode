package leetcode

import "testing"

var isPalindromeTests = []struct{
	Str string
	Result bool
}{
	{
		"race a car",
		false,
	},
	{
		"race e car",
		true,
	},
	{
		"!!!!!!",
		true,
	},
	{
		"A man, a plan, a canal: Panama",
		true,
	},
	{
		"0p",
		false,
	},
}

func TestIsPalindrome(t *testing.T) {
	for _, test := range isPalindromeTests {
		re := isPalindrome(test.Str)

		if re != test.Result {
			t.Errorf("failed! %v, act %v", test, re)
		}
	}
}
