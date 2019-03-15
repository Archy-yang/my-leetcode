package leetcode

func myPow(x float64, n int) float64 {
	if n < 0 {
		return 1.0 / myPow(x, -n)
	}

	var dfs func(float64, float64,int) float64
	dfs = func (r, x float64, n int) float64 {
		if n == 0 {
			return 1
		}

		r = dfs(r, x, n / 2)
		r = r * r
		if n & 1 == 1 {
			r = r * x
		}

		return r
	}

	return dfs(x, x, n)
}

func myPow1(x float64, n int) float64 {
	if n < 0 {
		return 1.0 / myPow1(x, -n)
	}

	r := 1.0
	p := x
	for {
		if n == 0 {
			break
		}
		if (n & 1) == 1 {
			r *= p
		}
		n = n >> 1
		p *= p
	}
	return r
}