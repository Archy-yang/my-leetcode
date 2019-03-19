package leetcode

import "testing"

var uniquePathsTests = []struct{
	inM, inN, out int
} {
	{
		3, 2, 3,
	},
	{
		7, 3, 28,
	},
}
func TestUniquePaths(t *testing.T)  {
	for _, test := range uniquePathsTests {
		re := uniquePaths(test.inM, test.inN);

		if re != test.out {
			t.Fatalf("faild. act: %v, exp: %v", re, test.out)
		}
	}
}
