package cn

import (
	"fmt"
	"testing"
)

func TestDeserialize(t *testing.T) {
	a := "[[123],123]"
	b := deserialize(a)
	fmt.Printf("%v", b)
}
