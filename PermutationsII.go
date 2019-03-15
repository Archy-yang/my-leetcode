package leetcode

import "sort"

func permuteUniqueIi(nums []int) [][]int {
	nLen := len(nums)
	if nLen < 1 {
		return [][]int{}
	}
	re := make([][]int, 0)
	used := make([]int, nLen)

	sort.Ints(nums)
	var dfs func([]int)
	dfs = func(one []int) {
		if len(one) == nLen {
			re = append(re, one)
			return
		}
		for i := 0; i < nLen; i++ {
			if used[i] == 1 {
				continue
			}
			if i > 0 && nums[i-1] == nums[i] && used[i-1] != 1 {
				continue
			}
			if used[i] == 0 {
				used[i] = 1
				t := make([]int, len(one))
				copy(t, one)
				t = append(t, nums[i])
				dfs(t)
				used[i] = 0
			}
		}
	}
	dfs([]int{})

	return re
}
