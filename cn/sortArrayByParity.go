package cn

func sortArrayByParity(nums []int) []int {
	i := 0
	j := len(nums) - 1

	for i < j {
		if nums[i] % 2 == 0 {
			i++
			continue
		}
		if nums[j] % 2 == 1 {
			j--
			continue
		}
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
	return nums
}
