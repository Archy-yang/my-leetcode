package cn

func lexicalOrder(n int) []int {
	s := make([]int, n)
	s[0] = 1
	for i := 1; i < n; i++ {
		if s[i-1]*10 <= n {
			s[i] = s[i-1] * 10
			continue
		}
		if s[i-1]%10 == 9 {
			s[i] = (s[i-1] + 1) / 10
			for ;s[i] % 10 == 0; {
				s[i] = s[i] / 10
			}
			continue
		}

		if s[i-1]+1 <= n {
			s[i] = s[i-1] + 1
			continue
		}
		if s[i-1]/10+1 <= n {
			s[i] = s[i-1]/10 + 1
			for ;s[i] % 10 == 0; {
				s[i] = s[i] / 10
			}
			continue
		}
	}

	return s
}
