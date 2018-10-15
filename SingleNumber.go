package leetcode


/**
参考了讨论，才找到解决办法。考的是位运算。还是需要加强位运算
 */
func singleNumber(nums []int) int {
	x := 0
	for _, num := range nums {
		x ^= num
	}

	return x
}