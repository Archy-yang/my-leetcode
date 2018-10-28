package leetcode

import (
	"testing"
	"fmt"
)

func TestOne(t *testing.T) {
	n := 1
	for i := 1 ; i <= 25; i++ {
		n *= i
	}
	fmt.Println(n)
}

func TestTranIntListToNode(t *testing.T) {
	re := tranIntListToNode([]int{1,2,3,4})

	for re != nil {
		fmt.Println(re)
		re = re.Next
	}
}
