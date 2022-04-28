package cn

func projectionArea(grid [][]int) int {
	s := 0
	is := make([]int, len(grid))
	js := make([]int, len(grid[0]))
	for i, v := range grid {
		for j, jv := range v {
			if jv != 0 {
				s++
			}
			if is[i] < jv {
				is[i] = jv
			}
			if js[j] < jv {
				js[j] = jv
			}
		}
	}
	for _, v := range is {
		s += v
	}
	for _, v := range js {
		s += v
	}

	return s
}

func projectionArea2(grid [][]int) int {
	s := 0
	for i := 0; i < len(grid); i++ {
		is := 0
		js := 0
		for j := 0; j < len(grid); j++ {
			if is < grid[i][j] {
				is = grid[i][j]
			}
			if js < grid[j][i] {
				js = grid[j][i]
			}
			if grid[i][j] > 0 {
				s++
			}
		}
		s = s + is + js
	}

	return s
}
