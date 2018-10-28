package leetcode

import "testing"

var countPrimesTests = []struct{
	input, output int
} {
	{
		1,0,
	},
	{
		2,0,
	},
	{
		3,1,
	},
	{
		10,4,
	},
}

func TestCountPrimes(t *testing.T) {
	for _, test := range countPrimesTests {
		re := countPrimes(test.input)

		if re != test.output {
			t.Fatalf("failed! %v %v", test.output, re)
		}
	}

}
func TestCountPrimes1(t *testing.T) {
	for _, test := range countPrimesTests {
		re := countPrimes1(test.input)

		if re != test.output {
			t.Fatalf("failed! %v %v", test.output, re)
		}
	}

}
