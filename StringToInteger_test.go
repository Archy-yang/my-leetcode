package leetcode

import (
	"testing"
)

var myAtoiTests = []struct {
	in  string
	out int
}{
	//{
	//	"4193 with words",
	//	4193,
	//},
	//{
	//	"3.4412",
	//	3,
	//},
	//{
	//	"1",
	//	1,
	//},
	//{
	//	" ",
	//	0,
	//},
	//{
	//	"-",
	//	0,
	//},
	{
		"9223372036854775808",
		2147483647,
	},
}

func TestMyAtoi(t *testing.T) {
	for _, test := range myAtoiTests {
		re := myAtoi(test.in)

		if re != test.out {
			t.Fatalf("fault, exp %v, act %v", test.out, re)
		}
	}
}
