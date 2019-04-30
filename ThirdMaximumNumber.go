package leetcode

func thirdMax(nums []int) int {
	l := []int{}
	for _, n := range nums {
		if len(l) == 0 {
			l = append(l, n)
			continue
		}
		if n == l[0] {
			continue
		}
		if n > l[0] {
			n, l[0] = l[0], n
		}
		if len(l) == 1{
			l = append(l, n)
			continue
		}
		if n == l[1] {
			continue
		}
		if n > l[1] {
			n, l[1] = l[1], n
		}
		if len(l) == 2 {
			l = append(l, n)
			continue
		}
		if n == l[2] {
			continue
		}
		if n > l[2] {
			n, l[2] = l[2], n
		}
	}

	if len(l) == 3 {
		return l[2]
	}
	return l[0]
}
