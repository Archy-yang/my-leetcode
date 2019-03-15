package leetcode

import (
	"testing"
	"fmt"
)

func TestPermuteUniqueIi(t *testing.T) {
	re := permuteUniqueIi([]int{1,1,2})
	//re := permuteUniqueIi([]int{1,1})
	//re := permuteUniqueIi([]int{1,2,3})

	fmt.Println(re)
}
