package cn

import "sort"

func heightChecker(heights []int) int {
	e := make([]int, len(heights))
	copy(e, heights)
	sort.Ints(e)
	r := 0
	for i, h := range e {
		if heights[i] != h {
			r ++
		}
	}
	return r
}
