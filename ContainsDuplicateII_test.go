package leetcode

import "testing"

var containsNearbyDuplicateTests = []struct{
	inputList []int
	inputK	int
	output bool
} {
	{
		[]int{1,2,3,1},
		3,
		true,
	},
	{
		[]int{1,2,3,1},
		2,
		false,
	},
}

func TestContainsNearbyDuplicate(t *testing.T) {
	for _, test := range containsNearbyDuplicateTests {
		re := containsNearbyDuplicate(test.inputList, test.inputK)
		if re != test.output {
			t.Fatalf("failed!! %v %v", test, re)
		}
	}
}
