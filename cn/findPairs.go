package cn

import "sort"

func findPairs(nums []int, k int) int {
	n := len(nums)
	sort.Ints(nums)
	a := 0
	for i := 0; i < n; i ++ {
		if i + 1 < n && nums[i] == nums[i+1] {
			continue
		}
		t := nums[i] + k
		if t == nums[i] {
			if find(nums[:i], t) {
				a++
			}
		} else {
			if find(nums[i+1:], t) {
				a++
			}
		}
	}

	return a
}

func find(nums []int, t int) bool {
	l := len(nums)
	if l < 1 {
		return false
	}
	h := l / 2
	if nums[h] == t {
		return true
	}
	if nums[h] > t {
		return find(nums[0:h], t)
	}
	return find(nums[h + 1:], t)
}
