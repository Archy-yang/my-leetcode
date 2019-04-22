package leetcode

func searchRange(nums []int, target int) []int {
	l := len(nums)
	if l == 0 || target > nums[l-1] || target < nums[0] {
		return []int{-1, -1}
	}
	lo := searchT(nums, target)
	if lo < 0 {
		return []int{-1, -1}
	}

	min, max := lo, lo
	for {
		stop := true
		if min-1 >= 0 && nums[min-1] == target {
			stop = false
			min--
		}
		if max+1 < l && nums[max+1] == target {
			stop = false
			max++
		}
		if stop {
			break
		}
	}

	return []int{min, max}
}

func searchT (nums []int, target int) int {
	start := 0
	end := len(nums)
	mid := 0
	for start <= end {
		mid = (start + end) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] < target {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}

	return -1
}
