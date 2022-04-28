package cn

import (
	"fmt"
	"testing"
)

func TestSortedArrayToBST(t *testing.T) {
	a := []int{-10,-3,0,5,9}
	b := sortedArrayToBST(a)
	fmt.Println(b)
}
