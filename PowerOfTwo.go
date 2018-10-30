package leetcode

func isPowerOfTwo(n int) bool {
	if n == 0 {
		return false
	}
	for {
		if n == 1 {
			return true
		}
		if n % 2 != 0 {
			return false
		}
		n = n / 2
	}
}

func isPowerOfTwo1(n int) bool {
	return (n > 0) &&  (n & (n-1) == 0)
}
