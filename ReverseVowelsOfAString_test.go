package leetcode

import "testing"

var reverseVowelsTests = []struct{
	in string
	out string
}{
	{"hello", "holle"},
	{"leetcode", "leotcede"},
}
func TestReverseVowels(t *testing.T)  {
	for _, test := range reverseVowelsTests {
		re := reverseVowels(test.in)

		if re != test.out {
			t.Fatalf("failed! exp: %v, ect: %v", test.out, re)
		}
	}
}
