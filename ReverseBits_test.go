package leetcode

import (
	"testing"
)

var reverseBitsTests = []struct{
	in, out uint32
} {
	{
		43261596,964176192,
	},
	{
		4294967293,3221225471,
	},
}

func TestReverseBits(t *testing.T)  {
	for _, test := range reverseBitsTests {
		re := reverseBits(test.in)

		if re != test.out {
			t.Fatalf("failed. exp: %v, act: %v", test.out, re)
		}
	}
}
