package leetcode

func rotate(nums []int, k int) {
	nLen := len(nums)
	k = k % nLen

	for i := 0; i < (nLen-k)/2; i++ {
		nums[i], nums[nLen-k-i-1] = nums[nLen-k-i-1], nums[i]
	}
	for i := 0; i < k/2; i++ {
		nums[nLen-k+i], nums[nLen-1-i] = nums[nLen-1-i], nums[nLen-k+i]
	}
	for i := 0; i < nLen/2; i++ {
		nums[i], nums[nLen-i-1] = nums[nLen-i-1], nums[i]
	}
}

