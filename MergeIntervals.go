package leetcode

import "sort"

func mergeIntervals(intervals [][]int) [][]int {
	if len(intervals) < 1 {
		return [][]int{}
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	re := [][]int{}
	one := intervals[0]
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] > one[1] {
			re = append(re, one)
			one = intervals[i]
		}
		if intervals[i][0] < one[0] {
			one[0] = intervals[i][0]
		}
		if intervals[i][1] > one[1] {
			one[1] = intervals[i][1]
		}
	}
	re = append(re, one)

	return re
}

func mergeIntervals2(intervals [][]int) [][]int {
	if len(intervals) < 1 {
		return [][]int{}
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	startMin, endMax := intervals[0][0], intervals[0][1]
	re := [][]int{}
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] > endMax {
			re = append(re, []int{startMin, endMax})
			startMin, endMax = intervals[i][0],intervals[i][1]
		}
		if intervals[i][0] < startMin {
			startMin = intervals[i][0]
		}
		if intervals[i][1] > endMax {
			endMax = intervals[i][1]
		}
	}
	re = append(re, []int{startMin, endMax})

	return re
}