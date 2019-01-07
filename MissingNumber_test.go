package leetcode

import "testing"

var missingNumberTests = []struct {
	in []int
	out int
}{
	//{
	//	in: []int{3,0,1},
	//	out: 2,
	//},
	//{
	//	in: []int{9,6,4,2,3,5,7,0,1},
	//	out: 8,
	//},
	//{
	//	in: []int{1},
	//	out: 0,
	//},
	//{
	//	in: []int{0},
	//	out: 1,
	//},
	//{
	//	in: []int{0, 1},
	//	out: 2,
	//},
	{
		in: []int{19,58,31,51,54,47,68,25,85,9,83,70,24,75,30,78,62,38,41,21,56,60,94,1,45,15,72,52,28,93,14,96,35,17,95,89,74,46,13,82,57,76,55,20,36,63,44,61,6,92,65,50,91,42,98,34,8,33,40,12,7,48,11,80,10,71,97,39,73,26,99,43,90,5,3,2,23,29,0,79,53,64,4,27,37,84,69,81,22,86,67,32,66,18,16,77,59,49,87},
		out: 88,
	},
}

func TestMissingNumber(t *testing.T) {
	for _, test := range missingNumberTests {
		re := missingNumber(test.in)

		if re != test.out {
			t.Errorf("failed!! exp: %+v, act: %+v", test.out, re)
		}
	}

}
