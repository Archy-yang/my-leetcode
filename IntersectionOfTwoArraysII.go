package leetcode

func intersect(nums1 []int, nums2 []int) []int {
	m := make(map[int]int)
	n := []int{}
	for _, num := range nums1 {
		if _, ok := m[num]; ok {
			m[num]++
		} else {
			m[num] = 1
		}
	}
	for _, num := range nums2 {
		if _, ok := m[num]; ok {
			if m[num] > 0 {
				n = append(n, num)
				m[num]--
			}
		}
	}
	return n
}