package cn

func isBoomerang(points [][]int) bool {
	v0 := points[0]
	v1 := points[1]
	v2 := points[2]

	if (v0[0] == v1[0] && v0[1] == v1[1]) ||
		(v0[0] == v2[0] && v0[1] == v2[1])  ||
		(v1[0] == v2[0] && v1[1] == v2[1])	{
		return false
	}

	x01 := v1[0] - v0[0]
	y01 := v1[1] - v0[1]
	var t01 float64 = 0
	if x01 == 0 {
		t01 = 1
	}else if y01 != 0 {
		t01 = float64(x01) / float64(y01)
	}

	x12 := v2[0] - v1[0]
	y12 := v2[1] - v1[1]
	var t21 float64 = 0
	if x12 == 0 {
		t21 = 1
	} else if y12 != 0 {
		t21 = float64(x12) / float64(y12)
	}
	if t01 != t21 {
		return true
	}
	return false
}
