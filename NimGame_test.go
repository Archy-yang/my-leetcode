package leetcode

import "testing"

var canWinNimTests = []struct{
	in int
	out bool
}{
	{1, true},
	{2, true},
	{3, true},
	{4, false},
	{5, true},
	{6, true},
	{7, true},
}

func TestCanWinNim(t *testing.T) {
	for _, test := range canWinNimTests {
		re := canWinNim(test.in)

		if re != test.out {
			t.Fatalf("failed. exp: %v, act: %v", test.out, re)
		}
	}
}
