package leetcode

import (
	"sort"
)

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)

	all := [][]int{}
	var dfs func([]int, int, []int)
	dfs = func(candidates []int, target int, q []int) {
		if len(candidates) < 1 {
			return
		}
		for i := range candidates {
			if candidates[i] > target {
				return
			}
			l := len(q)
			q = append(q, candidates[i])
			if candidates[i] == target {
				t := make([]int, l+1)
				copy(t, q)
				all = append(all, t)
				return
			}
			dfs(candidates[i:], target - candidates[i], q)
			q = q[0:l]
		}

	}
	dfs(candidates, target, []int{})

	return all
}
