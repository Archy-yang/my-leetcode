package leetcode

import (
	"testing"
	"fmt"
)

func TestOne(t *testing.T) {
	n := 1
	for i := 1 ; i <= 25; i++ {
		n *= i
	}
	fmt.Println(n)
}
