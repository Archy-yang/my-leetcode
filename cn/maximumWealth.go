package cn

func maximumWealth(accounts [][]int) int {
	max := 0
	for _, account := range accounts {
		tmp := 0
		for _, bank := range account {
			tmp = tmp + bank
		}
		if (max < tmp) {
			max = tmp
		}
	}

	return max
}
