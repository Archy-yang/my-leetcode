package leetcode

import (
	"testing"
)

func TestMerge(t *testing.T)  {
	as := []struct{
		nums1 []int
		m int
		nums2 []int
		n int
		re []int
	}{
		{
			nums1:[]int{1,2,3,0,0,0},
			m: 3,
			nums2: []int{2,5,6},
			n:3,
			re: []int{1,2,2,3,5,6},
		},
		{
			nums1:[]int{1,2,3,0,0,0},
			m: 0,
			nums2: []int{2,5,6},
			n:3,
			re: []int{2,5,6},
		},
		{
			nums1:[]int{1,2,3,0,0,0},
			m: 3,
			nums2: []int{2,5,6},
			n:0,
			re: []int{1,2,3},
		},
		{
			nums1:[]int{0},
			m: 0,
			nums2: []int{1},
			n:1,
			re: []int{1},
		},
	}
	for _, a := range as {
		re := merge(a.nums1, a.m, a.nums2, a.n)
		if len(re) != len(a.re) {

		}
		for i, v := range re {
			if v != a.re[i] {
				t.Errorf("%q result err, expcet %v, act %v", a, a.re, re)
			}
		}
	}
}
