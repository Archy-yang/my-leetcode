package cn

import (
	"strconv"
	"unicode"
)

// This is the interface that allows for creating nested lists.
// You should not implement it, or speculate about its implementation
type NestedInteger struct {
	i int
	l []*NestedInteger
}

// Return true if this NestedInteger holds a single integer, rather than a nested list.
func (n NestedInteger) IsInteger() bool {
	return len(n.l) == 0
}

// Return the single integer that this NestedInteger holds, if it holds a single integer
// The result is undefined if this NestedInteger holds a nested list
// So before calling this method, you should have a check
func (n NestedInteger) GetInteger() int {
	return n.i
}

// Set this NestedInteger to hold a single integer.
func (n *NestedInteger) SetInteger(value int) {
	n.i = value
}

// Set this NestedInteger to hold a nested list and adds a nested integer to it.
func (n *NestedInteger) Add(elem NestedInteger) {
	n.l = append(n.l, &elem)
}

// Return the nested list that this NestedInteger holds, if it holds a nested list
// The list length is zero if this NestedInteger holds a single integer
// You can access NestedInteger's List element directly if you want to modify it
func (n NestedInteger) GetList() []*NestedInteger {
	return n.l
}

func deserialize(s string) *NestedInteger {
	p := 0
	l := len(s)
	return f(s, &p, l)
}

func f(s string, p *int, l int) *NestedInteger {
	if *p >= l || s[*p] == ']' {
		return nil
	}
	n := &NestedInteger{}
	si := ""
	for ; *p < l && s[*p] != ']' && s[*p] != '[' && s[*p] != ','; *p++ {
		si = si + string(s[*p])
	}
	if len(si) > 0 {
		v, _ := strconv.Atoi(si)
		n.SetInteger(v)
		return n
	}
	for ; *p < l && (s[*p] == '[' || s[*p] == ','); {
		*p++
		t := f(s, p, l)
		n.Add(*t)
		if s[*p] == ']' {
			*p++
			break
		}
	}

	return n
}


func deserializeV1(s string) *NestedInteger {
	if s[0] != '[' {
		num, _ := strconv.Atoi(s)
		ni := &NestedInteger{}
		ni.SetInteger(num)
		return ni
	}
	stack, num, negative := []*NestedInteger{}, 0, false
	for i, ch := range s {
		if ch == '-' {
			negative = true
		} else if unicode.IsDigit(ch) {
			num = num*10 + int(ch-'0')
		} else if ch == '[' {
			stack = append(stack, &NestedInteger{})
		} else if ch == ',' || ch == ']' {
			if unicode.IsDigit(rune(s[i-1])) {
				if negative {
					num = -num
				}
				ni := NestedInteger{}
				ni.SetInteger(num)
				stack[len(stack)-1].Add(ni)
			}
			num, negative = 0, false
			if ch == ']' && len(stack) > 1 {
				stack[len(stack)-2].Add(*stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
		}
	}
	return stack[len(stack)-1]
}
