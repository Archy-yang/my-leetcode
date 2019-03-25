package leetcode

func letterCombinations(digits string) []string {
	m := []string{"abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
	re := []string{}
	for i := 0; i < len(digits); i ++ {
		a := digits[i] - '2'
		c := m[a]
		tmp := []string{}
		if len(re) == 0 {
			for j := 0; j < len(c); j++ {
				tmp = append(tmp, string(c[j]))
			}
		} else {
			for _, r := range re {
				for j := 0; j < len(c); j++ {
					r := r + string(c[j])
					tmp = append(tmp, r)
				}
			}
		}
		re = tmp
	}

	return re
}
