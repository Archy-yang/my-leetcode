package leetcode

import (
	"testing"
	"fmt"
)

func TestSolveNQueens(t *testing.T) {
	re := solveNQueens(7);
	re1 := solveNQueens1(7);


	for i, _ := range re {
		for j, _ := range re[i] {
			if re[i][j] != re1[i][j] {
				fmt.Println(re[i][j], re1[i][j])
				fmt.Println(i, j)
			}
		}
	}
}
