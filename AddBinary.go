package leetcode

import "strconv"

func addBinary(a string, b string) string {
	al := len(a)
	bl := len(b)
	if al == 0 {
		return b
	}
	if bl == 0 {
		return a
	}

	var ls, ss string = a, b
	var ll, sl int = al, bl
	if (al < bl) {
		ls, ss = b, a
		ll, sl = bl, al
	}
	var f int
	var r string
	for  ll > 0 {
		li := int(ls[ll-1] - 48)
		var si int
		if sl - 1 >= 0 {
			si = int(ss[sl-1] - 48)
		} else {
			si = 0
		}

		n := li+si+f
		if n < 2 {
			f = 0
		} else {
			f = 1
			n = n - 2
		}
		r = strconv.Itoa(n) + r

		sl--
		ll--
	}
	if f > 0 {
		r = strconv.Itoa(f) + r
	}

	return r
}