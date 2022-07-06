package cn

import (
	"fmt"
	"testing"
)

func TestDuplicateZeros(t *testing.T) {
	a := []int{0,0,0,0,0,0,0}
	//a := []int{8,4,5,0,0,0,0,7}
	duplicateZeros(a)
	fmt.Println(a)
}
