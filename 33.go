package leetcode

func search33(nums []int, target int) int {
	if len(nums) < 1 {
		return -1
	}
	if len(nums) == 1 && nums[0] != target{
		return -1
	}
	minP := findMin(nums)
	if target >= nums[0] && minP > 0{
		return searchNum(nums[0:minP], target)
	} else {
		t := searchNum(nums[minP:], target)
		if t < 0 {
			return t
		}
		return t+minP
	}
}

func findMin(nums []int) int {
	start := 0
	end := len(nums) - 1
	if end == 0 || nums[end] > nums[start]{
		return 0
	}

	for end > start {
		mid := (start + end + 1) / 2
		if nums[mid] < nums[mid-1] {
			return mid
		}
		if nums[mid] > nums[end] {
			start = mid
		} else {
			end = mid
		}
	}
	return 0
}

func searchNum(nums []int, target int) int {
	start := 0
	end := len(nums) - 1

	if target > nums[end] || target < nums[start] {
		return -1
	}

	for end >= start {
		mid := (start + end) / 2

		if nums[mid] == target {
			return mid
		}
		if nums[mid] > target {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}

	return -1
}

