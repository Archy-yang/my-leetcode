package leetcode

func containsNearbyDuplicate(nums []int, k int) bool {
	m := map[int]int{}
	for i, num := range nums {
		if index, ok := m[num]; ok &&  (i-index) <= k {
			return true
		}
		m[num] = i
	}
	return false
}
