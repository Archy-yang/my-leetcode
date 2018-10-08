package leetcode

import (
	"strconv"
)

func countAndSay(n int) string {
	a := ""
	for i := 0; i < n; i++ {
		var b string
		la := len(a)

		c := 1
		for j := 0; j < la-1; j++ {
			if a[j] == a[j+1] {
				c++
				continue
			}

			b += strconv.Itoa(c) + string(a[j])
			c = 1
		}

		if la > 0 {
			a = b + strconv.Itoa(c) + string(a[la-1])
		} else {
			a = b + strconv.Itoa(c)
		}
	}

	return a
}
