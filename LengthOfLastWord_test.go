package leetcode

import (
	"testing"
	"fmt"
)

var tests = map[string]int {
	"hello world": 5,
	"a ": 1,
	" a": 1,
}

func Test_LengthOfLastWord(t *testing.T)  {
	for word, num := range tests {
		act := lengthOfLastWord(word)

		if (act != num) {
			fmt.Printf("error : %v, expect: %v, act %v\n", word, num, act)
		} else {
			fmt.Printf("success : %v\n", word)
		}
	}
}
