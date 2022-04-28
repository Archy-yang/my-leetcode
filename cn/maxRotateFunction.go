package cn

func maxRotateFunction(nums []int) int {
	nl := len(nums)
	f := 0
	s := 0
	for i, n := range nums {
		f += i * n
		s += n
	}
	max := f
	for i := 1; i < nl; i++ {
		f = f + s - nl * nums[nl-i]
		if max < f {
			max = f
		}
	}

	return max
}
