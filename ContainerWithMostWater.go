package leetcode

func maxArea(height []int) int {
	max := 0
	for i := 0; i < len(height) - 1; i ++ {
		for j := i + 1; j < len(height); j++ {
			min := height[i]
			if min > height[j] {
				min = height[j]
			}
			if max < (j - i) * min {
				max = (j-i) * min
			}
		}
	}

	return max
}