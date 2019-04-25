package leetcode

func trap(height []int) int {
	if len(height) < 1 {
		return 0
	}
	sum := 0
	max := height[0]

	i := 0
	for {
		tmp := 0
		tmp2 := 0
		f := false
		for _, n := range height {
			if max < n {
				max = n
			}
			if n > i {
				tmp += tmp2
				tmp2 = 0
				f = true
			} else {
				if f {
					tmp2++
				}
			}
		}

		if max == i {
			break
		}
		sum += tmp
		i++
	}
	return sum
}

func trap2(height []int) int {
	left := 0
	right := len(height) - 1
	w := 0
	l := 0
	for left < right {
		if height[left] < height[right] {
			if l < height[left] {
				l = height[left]
			}
			w += l - height[left]
			left++
		} else {
			if l < height[right] {
				l = height[right]
			}
			w += l - height[right]
			right--
		}
	}
	return w
}
