package leetcode

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	m := map[int]int{}
	for i:=0; i < len(s); i++ {
		is := int(s[i])
		n := int(s[i])
		if (is & 192 == 192) {
			i++
			n = n*1000 + int(s[i])
		}
		if (is & 224 == 224) {
			i++
			n = n*1000 + int(s[i])
		}
		if (is & 240 == 240) {
			i++
			n = n*1000 + int(s[i])
		}
		m[n]++
	}
	s = t
	for i:=0; i < len(s); i++ {
		is := int(s[i])
		n := int(s[i])
		if (is & 192 == 192) {
			i++
			n = n*1000 + int(s[i])
		}
		if (is & 224 == 224) {
			i++
			n = n*1000 + int(s[i])
		}
		if (is & 240 == 240) {
			i++
			n = n*1000 + int(s[i])
		}
		m[n]--
	}

	for _, num := range m {
		if num != 0 {
			return false
		}
	}

	return true
}
