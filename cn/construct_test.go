package cn

import (
	"fmt"
	"testing"
)

func TestConstruct(t *testing.T) {
	a := [][]int{{1,1,1,1,0,0,0,0},{1,1,1,1,0,0,0,0},{1,1,1,1,1,1,1,1},{1,1,1,1,1,1,1,1},{1,1,1,1,0,0,0,0},{1,1,1,1,0,0,0,0},{1,1,1,1,0,0,0,0},{1,1,1,1,0,0,0,0}}
	b := construct(a)
	fmt.Printf("%v", b)
}

