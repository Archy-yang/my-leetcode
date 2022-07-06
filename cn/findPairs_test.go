package cn

import (
	"fmt"
	"testing"
)

func TestFindPairs(t *testing.T) {
	a := []int {3,1,4,1,5}
	//a := []int {1,1,1,1}
	k := 2
	b := findPairs(a, k)
	fmt.Println(b)
}

