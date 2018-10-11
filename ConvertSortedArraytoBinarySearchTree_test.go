package leetcode

import (
	"testing"
	"fmt"
)

func TestSortedArrayToBST(t *testing.T) {
	r := sortedArrayToBST([]int{-10,-3,0,5,9})
	l := levelOrderBottom(r)
	fmt.Println(l)
}

