package leetcode

func nthUglyNumber(n int) int {
	if n == 0 || n == 1{
		return n
	}
	a := make([]int, n)
	a[0] = 1
	l := [3]int{0,0,0}
	num := 0
	for i := 1; i < n; i++ {
		l2 := a[l[0]] * 2
		l3 := a[l[1]] * 3
		l5 := a[l[2]] * 5
		next := l2
		if next > l3 {
			next = l3
		}
		if next > l5 {
			next = l5
		}
		if next == l2 {
			l[0]++
		}
		if next == l3 {
			l[1]++
		}
		if next == l5 {
			l[2]++
		}
		a[i] = next
		num = next
	}

	return num
}
