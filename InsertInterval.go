package leetcode

func insertInterval(intervals [][]int, newInterval []int) [][]int {
	if len(newInterval) < 1 {
		return intervals
	}
	intervals = append(intervals, newInterval)
	re := [][]int{}
	f := false
	for _, interval := range intervals {
		if interval[0] > newInterval[1] {
			if !f {
				re = append(re, newInterval)
				f = true
			}
			re = append(re, interval)
			continue
		}
		if interval[1] < newInterval[0] {
			re = append(re, interval)
			continue
		}

		if interval[0] <= newInterval[0] && interval[1] >= newInterval[0] {
			newInterval[0] = interval[0]
		}
		if interval[0] <= newInterval[1] && interval[1] >= newInterval[1] {
			newInterval[1] = interval[1]
		}
	}
	if !f {
		re = append(re, newInterval)
	}

	return re
}
