package cn

import (
	"fmt"
	"testing"
)

func TestProjectionArea(t *testing.T) {
	a := [][]int{{1,2}, {3,4}}
	b := projectionArea2(a)
	fmt.Println(b)
}
