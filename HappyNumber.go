package leetcode

func isHappy(n int) bool {
	nums := map[int]bool{}

	for {
		if _, ok := nums[n]; ok {
			return false
		}
		nums[n] = true

		var t int
		for {
			s := n%10
			t += s * s
			if n < 10 {
				break
			}
			n = n / 10
		}
		if t == 1 {
			return true
		}
		n = t
	}
}
