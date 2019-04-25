package leetcode

import (
	"sort"
)

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)

	all := [][]int{}
	var dfs func([]int, int, []int)
	dfs = func(candidates []int, target int, q []int) {
		if len(candidates) < 1 {
			return
		}
		lc := len(candidates)
		for i := 0; i < lc; i ++ {
			if candidates[i] > target {
				return
			}
			l := len(q)
			t := candidates[i]
			q = append(q, candidates[i])
			if candidates[i] == target {
				t := make([]int, l+1)
				copy(t, q)
				all = append(all, t)
				return
			}
			dfs(candidates[i+1:], target - candidates[i], q)
			for i + 1 < lc && candidates[i + 1] == t {
				i++
			}
			q = q[0:l]
		}

	}
	dfs(candidates, target, []int{})

	return all
}
