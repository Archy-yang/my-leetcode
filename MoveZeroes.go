package leetcode

func moveZeroes(nums []int)  {
	l := len(nums)

	for i := 0; i < l; {
		if nums[i] == 0 {
			for j := i; j < l-1; j++ {
				nums[j] = nums[j+1]
				nums[j+1] = 0
			}
			l--
		} else {
			i++
		}
	}
}
