package leetcode
/**
此题参考以后才找到答案。推倒数学模型。
需考虑，不能只看每一个nums，要看当前num与之前已累加的数的比较。解题模型如下：
假设b[i]是前i个数中最大的值，那么第i+1个数的最大值为：
b[i+1] = max(b[i], nums[i+1]+b[i-1]),那么，对于第n个数：
b[n] = max(b[n-1], nums[n]+b[n-2])。
 */
func rob(nums []int) int {
	nLen := len(nums)

	if nLen == 0 {
		return 0
	}
	if nLen == 1 {
		return nums[0]
	}
	if nums[1] < nums[0] {
		nums[1] = nums[0]
	}
	for i:= 2; i< nLen ; i++ {
		nums[i] = nums[i] + nums[i-2]
		if nums[i] < nums[i-1] {
			nums[i] = nums[i-1]
		}
	}

	return nums[nLen-1]
}