package leetcode

import "fmt"

func countPrimes(n int) int {
	t := 0
	for i:=2; i<n; i++ {
		if isPrimes(i) {
			t++
		}
	}

	return t
}

func isPrimes(n int) bool {
	if n == 3 || n == 2 {
		return true
	}
	for i:=2; i*i <= n; i++ {
		if n % i == 0 {
			return false
		}
	}
	return true
}

func countPrimes1(n int) int {
	m := make(map[int]bool, n)

	if n < 3 {
		return 0
	}
	if n == 3 {
		return 1
	}

	for i:=2; i*i < n; i++ {
		if !m[i] {
			for j := i*i; j < n; j += i {
				m[j] = true
			}
		}
	}
	fmt.Println(m)
	count := 0
	for i := 2; i < n; i++ {
		if v, ok := m[i]; !ok ||  !v {
			count++
		}
	}

	return count
}
