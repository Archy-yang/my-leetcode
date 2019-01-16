package leetcode

import "math"

func isPowerOfThree(n int) bool {
	if n == 0 {
		return false
	}
	p := int(math.Log2(float64(n)) / math.Log2(3))

	return math.Pow(3, float64(p)) == float64(n) || math.Pow(3, float64(p + 1)) == float64(n)
}
