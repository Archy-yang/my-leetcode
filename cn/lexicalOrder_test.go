package cn

import (
	"fmt"
	"testing"
)

func TestLexicalOrder(t *testing.T) {
	a := 192
	//a := 1003
	b := lexicalOrder(a)
	fmt.Println(b)
}
