package leetcode

import (
	"sort"
	"fmt"
)

func threeSum(nums []int) [][]int {
	re := [][]int{}
	if len(nums) < 3 {
		return re
	}
	sort.Ints(nums)
	mm := map[string]int{}
	for i := 0; i < len(nums) -2; i ++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue;
		}
		m := map[int]int{}
		for j := i + 1; j < len(nums); j ++ {
			if _, ok := m[nums[j]]; ok {
				k := fmt.Sprintf("%d_%d_%d", nums[i], -nums[i]-nums[j], nums[j])
				if _, ok := mm[k]; ok {
					continue
				}
				mm[k] = 1
				re = append(re, []int{nums[i], -nums[i]-nums[j], nums[j]})
			} else {
				m[-nums[i]-nums[j]] = 1
			}
		}
	}

	return re
}