package leetcode

func firstMissingPositive(nums []int) int {
	ln := len(nums)
	for i := 0 ; i < ln; i ++ {
		for {
			if nums[i] == i+1 {
				break
			}
			if nums[i] > ln || nums[i] < 1 {
				nums[i] = 0
				break
			}
			if nums[nums[i] - 1] == nums[i] {
				nums[i] = 0
				break
			}
			nums[i], nums[nums[i] - 1] = nums[nums[i] - 1], nums[i]
		}
	}

	for i := range nums {
		if nums[i] == 0 {
			return i+1
		}
	}
	return ln+1
}
