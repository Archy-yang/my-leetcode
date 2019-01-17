package leetcode

func intersection(nums1 []int, nums2 []int) []int {
	l := len(nums1)
	s := len(nums2)
	nums := []int{}
	if l < s {
		l, s = s, l
	}

	m := make(map[int]bool, l)
	for _, num := range nums1 {
		m[num] = false
	}
	for _, num := range nums2 {
		if v, ok := m[num]; ok {
			if v == false {
				nums = append(nums, num)
			}
			m[num] = true
		}
	}

	return nums
}