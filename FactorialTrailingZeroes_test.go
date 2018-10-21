package leetcode

import "testing"

var trailingZeroesTests = []struct{
	intput, output int
}{
	//{
	//	0,0,
	//},
	//{
	//	1,0,
	//},
	//{
	//	2, 0,
	//},
	//{
	//	5,1,
	//},
	//{
	//	7,1,
	//},
	//{
	//	10,2,
	//},
	{
		15,3,
	},
	{
		20,4,
	},
	{
		25,6,
	},
	{
		50,12,
	},
}

func TestTrailingZeroes(t *testing.T) {
	for _, test := range trailingZeroesTests {
		re := trailingZeroes(test.intput)

		if re != test.output {
			t.Fatalf("failed!  %v %v", test, re)
		}
	}
}
