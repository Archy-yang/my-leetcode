package leetcode

func merge(nums1 []int, m int, nums2 []int, n int)  []int {
	if m == 0 {
		return nums2[0:n]
	}
	if n == 0 {
		return nums1[0:m]
	}
	for i := m-1 ; i >= 0 ; i-- {
		nums1[n+i] = nums1[i]
	}
	i, j := 0, 0
	for {
		if j == n {
			break
		}
		if i-j >= m || nums1[n+i-j] > nums2[j] {
			nums1[i] = nums2[j]
			j++
		} else {
			nums1[i] = nums1[n+i-j]
		}

		i++
	}
	return nums1[0:n+m]
}
