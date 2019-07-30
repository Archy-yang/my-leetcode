package leetcode

import (
	"sort"
)

func groupAnagrams(strs []string) [][]string {
	if len(strs) < 1 {
		return [][]string{}
	}
	m := map[string][]string{}

	f := func (in []byte) []byte {
		sort.Slice(in, func(i, j int) bool{
			return in[i] < in[j]
		})

		return in
	}

	for _, str := range strs {
		t := string(f([]byte(str)))

		_, ok := m[t]
		if !ok {
			m[t] = []string{}
		}
		m[t] = append(m[t], str)
	}

	r := [][]string{}
	for _, s := range m {
		r = append(r, s)
	}

	return r
}
