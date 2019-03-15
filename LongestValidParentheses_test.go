package leetcode

import (
	"testing"
	"fmt"
)

func TestLongestValidParentheses(t *testing.T) {
	re := longestValidParentheses("(()")

	fmt.Println(re)
}
