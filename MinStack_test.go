package leetcode

import (
	"testing"
)

var minStackTests = []struct{
	t string
	v int
	r int
} {
	//{
	//	t: "push",
	//	v: -2,
	//},
	//{
	//	t: "push",
	//	v: 0,
	//},
	//{
	//	t: "push",
	//	v: -3,
	//},
	//{
	//	t: "getmin",
	//	r: -3,
	//},
	//{
	//	t: "pop",
	//	v: -2,
	//},
	//{
	//	t: "top",
	//	r: 0,
	//},
	//{
	//	t: "getmin",
	//	r: -2,
	//},
	{
		t: "push",
		v: 2147483646,
	},
	{
		t: "push",
		v: 2147483646,
	},
	{
		t: "push",
		v: 2147483647,
	},
	{
		t: "pop",
		v: 2147483647,
	},
	{
		t: "pop",
		v: 2147483647,
	},
	{
		t: "pop",
		v: 2147483647,
	},
	{
		t: "push",
		v: 2147483647,
	},
	{
		t: "getmin",
		r: 2147483647,
	},
}

func TestMinStack(t *testing.T) {

	stack := Constructor()
	for _, test := range  minStackTests {
		var re int
		if test.t == "push" {
			stack.Push(test.v)
		}
		if test.t == "pop" {
			stack.Pop()
		}
		if test.t == "top" {
			re = stack.Top()
		}
		if test.t =="getmin" {
			re = stack.GetMin()
		}
		if re != test.r {
			t.Fatalf("failed! %v %v %v", stack, test, re)
		}
	}
}
