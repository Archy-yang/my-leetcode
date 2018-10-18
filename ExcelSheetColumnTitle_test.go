package leetcode

import "testing"

var convertToTitleTests = []struct{
	input int
	output string
} {
	{
		1,
		"A",
	},
	{
		26,
		"Z",
	},
	{
		28,
		"AB",
	},
}
func TestConvertToTitle(t *testing.T) {
	for _, test := range convertToTitleTests {
		re := convertToTitle(test.input)

		if re != test.output {
			t.Fatalf("failed !  %v %v", test, re)
		}
	}
}

