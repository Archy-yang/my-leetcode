package cn

func duplicateZeros(arr []int)  {
	lp := len(arr)
	p := lp
	for i := 0; i < lp; i++ {
		if arr[i] == 0 {
			p++
		}
	}
	p--
	for i := lp - 1; i >= 0; i-- {
		if arr[i] == 0 {
			if p < lp {
				arr[p] = 0
			}
			p--
		}
		if p < lp {
			arr[p] = arr[i]
		}
		p--
	}
}
