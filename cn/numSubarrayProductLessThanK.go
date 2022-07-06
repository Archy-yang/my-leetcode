package cn

func numSubarrayProductLessThanK(nums []int, k int) int {
	if k < 1 {
		return 0
	}
	i, j, a, t, l := 0, 0, 0, 1, len(nums)
	for ; i < l && j < l; {
		for ; t * nums[j] >= k && i <= j; {
			t = t / nums[i]
			i++
		}
		if i <= j {
			t = t * nums[j]
			a += j - i + 1
		} else {
			t = 1
		}
		j++
	}

	return a
}
