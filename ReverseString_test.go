package leetcode

import "testing"

var reverseStringTests = []struct{
	in, out string
}{
	{
		in: "",
		out: "",
	},
	{
		in: "hello",
		out: "olleh",
	},
}
func TestReverseString(t *testing.T) {
	for _, test := range reverseStringTests {
		re := reverseString(test.in)

		if re != test.out {
			t.Fatalf("failed! act: %v, exp: %v", re, test.out)
		}
	}
}
func TestReverseString2(t *testing.T) {
	for _, test := range reverseStringTests {
		re := reverseString2(test.in)

		if re != test.out {
			t.Fatalf("failed! act: %v, exp: %v", re, test.out)
		}
	}
}
