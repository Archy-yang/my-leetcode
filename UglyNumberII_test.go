package leetcode

import "testing"

var nthUglyNumberTests = []struct{
	input int
	output int
} {
	//{
	//	0, 0,
	//},
	//{
	//	1, 1,
	//},
	//{
	//	2, 2,
	//},
	//{
	//	10, 12,
	//},
	{
		11, 15,
	},
}

func TestNthUglyNumber(t *testing.T) {
	for _, test := range nthUglyNumberTests {
		re := nthUglyNumber(test.input)

		if re != test.output {
			t.Fatalf("failed!! %v %v", test, re)
		}
	}
}
