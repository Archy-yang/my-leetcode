package cn

import (
	"fmt"
	"testing"
)

func TestAddBinary(t *testing.T) {
	a := "10"
	b := "1"
	c := addBinary(a, b)
	fmt.Println(c)
}