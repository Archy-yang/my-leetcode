package cn

import (
"fmt"
"testing"
)

func TestFindDiagonalOrder(t *testing.T) {
	b := findDiagonalOrder([][]int{{2,3}})
	fmt.Println(b)
}
