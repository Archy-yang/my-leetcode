package cn

import (
	"fmt"
	"testing"
)

func TestFindSubstringInWraproundString(t *testing.T) {
	a := "zab"
	b := findSubstringInWraproundString(a)
	fmt.Println(b)
}
