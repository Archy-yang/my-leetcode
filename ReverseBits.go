package leetcode

func reverseBits(num uint32) uint32 {
	s := uint32(0)
	u := uint32(1)
	for i := 0; i < 32; i ++ {
		t := num & u
		u = u << 1
		if i > 0 {
			s = s << 1
		}
		if t > 0 {
			s += 1
		}
	}

	return s
}


func reverseBits2(n uint32) uint32 {
	n = (n >> 16) | (n << 16)
	n = ((n & 0xff00ff00) >> 8) | ((n & 0x00ff00ff) << 8)
	n = ((n & 0xf0f0f0f0) >> 4) | ((n & 0x0f0f0f0f) << 4)
	n = ((n & 0xcccccccc) >> 2) | ((n & 0x33333333) << 2)
	n = ((n & 0xaaaaaaaa) >> 1) | ((n & 0x55555555) << 1)
	return n
}