package cn

import (
	"fmt"
	"testing"
)

func TestLengthOfLastWord(t *testing.T) {
	a := " ddd ddd dddd "
	b := lengthOfLastWord(a)
	fmt.Println(b)
}