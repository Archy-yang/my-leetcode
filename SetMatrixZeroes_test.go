package leetcode

import (
	"testing"
	"fmt"
)

func TestSetMatrixZeroes(t *testing.T) {
	//m := [][]int{
	//	{1,1,1},{0,1,2},
	//}
	//m := [][]int{
	//	{1,1,1},{1,0,1},{1,1,1},
	//}
	m := [][]int{
		{-4,-2147483648,6,-7,0},{-8,6,-8,-6,0},{2147483647,2,-9,-6,-10},
	}
	setZeroes(m)
	fmt.Println(m)
}
