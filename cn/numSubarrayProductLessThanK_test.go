package cn

import (
	"fmt"
	"testing"
)

func TestNumSubarrayProductLessThanK(t *testing.T) {
	a := []int{1,1,1}
	b := numSubarrayProductLessThanK(a, 1)
	fmt.Println(b)
}