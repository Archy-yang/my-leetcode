package cn

func maxProfit(prices []int) int {
	l := len(prices)
	if (l < 2) {
		return 0
	}
	max, min := prices[0], prices[0]
	profit := 0
	pt := 0

	for i := 1; i < l; i++ {
		if prices[i] < min {
			max, min = prices[i], prices[i]
			pt = 0
		}
		if prices[i] > max {
			pt += prices[i]-max
			max = prices[i]
		}
		if profit < pt {
			profit = pt
		}
	}

	return profit
}
