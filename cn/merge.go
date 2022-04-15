package cn

func merge(nums1 []int, m int, nums2 []int, n int)  {
	if n < 1 {
		return
	}
	allP := m + n - 1
	m--
	n--

	for ;allP >=0; allP-- {
		if m < 0 {
			nums1[allP] = nums2[n]
			n--
			continue
		}
		if n < 0 {
			break
		}
		if nums1[m] > nums2[n] {
			nums1[allP] = nums1[m]
			m--
		} else {
			nums1[allP] = nums2[n]
			n--
		}
	}

	return
}