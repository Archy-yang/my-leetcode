package leetcode

func canJump(nums []int) bool {
	if len(nums) < 1 {
		return true
	}
	maxIndex := nums[0]
	for i:=0; i <= maxIndex && i < len(nums);  i ++{
		if nums[i] + i > maxIndex {
			maxIndex = nums[i] + i
		}
	}
	if maxIndex < len(nums) -1 {
		return false
	}

	return true
}