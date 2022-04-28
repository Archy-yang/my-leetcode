package cn

import "math"

func shortestToChar(s string, c byte) []int {
	r := make([]int, len(s))
	t := []int{}

	for k, v := range s {
		if uint8(v) == c {
			t = append(t, k)
		}
	}
	for k, _ := range s {
		r[k] = k - t[0]
		for _, p := range t {
			if math.Abs(float64(k - p)) < float64(r[k]) {
				r[k] = int(math.Abs(float64(k - p)))
			}
		}
	}
	return r
}
