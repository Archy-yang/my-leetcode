package leetcode

func jump(nums []int) int {
	l := len(nums)
	step := make([]int, l)

	for i, n := range nums {
		for j := 1; j <= n; j++ {
			s := step[i] + 1
			if i + j >= l {
				break
			}
			if step[i + j] == 0 || step[i+j] > s {
				step[i+j] = s
			}
		}
	}
	return step[l-1]
}

//  贪心。。。 找每一步里能到的最大的步数
func jumpCankao(nums []int) int {
	max := 0
	s := 0
	e := 0

	for i := 0; i < len(nums) -1 ; i ++ {
	//for i, n := range nums {
		if max < nums[i]+i {
			max = nums[i] + i
		}
		if e == i {
			s++
			e = max
		}
	}

	return s
}
