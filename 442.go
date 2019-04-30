package leetcode

// Find All Duplicates in an Array
func findDuplicates(nums []int) []int {
	re := []int{}

	for i := 0; i < len(nums); i++ {
		for nums[i] != i+1 {
			if nums[nums[i] - 1] == nums[i] {
				break
			} else {
				nums[i], nums[nums[i] - 1] = nums[nums[i] - 1], nums[i]
			}
		}
	}
	for i := 0; i< len(nums); i++ {
		if nums[i] != i + 1 {
			re = append(re, nums[i])
		}
	}
	return re
}
