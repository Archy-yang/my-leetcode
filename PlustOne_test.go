package leetcode

import (
	"testing"
	"fmt"
)
type oneTest struct {
	input []int
	expect []int
}
var plusOneTests  = []oneTest {
	oneTest{
		input: []int{9,9},
		expect:[]int{1,0,0},
	},
}

func TestPlusOne(t *testing.T) {
	for _, oneTest := range plusOneTests {
		act := plusOne(oneTest.input)

		if (len(act) != len(oneTest.expect)) {
			fmt.Printf("error : %v, act %v\n", oneTest, act)
		} else {
			fmt.Printf("success : %v\n", oneTest)
		}
	}
}
