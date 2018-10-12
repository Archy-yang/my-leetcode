package leetcode

func maxProfitII(prices []int) int {
	profit := 0
	if len(prices) > 0 {
		minPrice := prices[0]
		lenth := len(prices)
		for index, price := range prices {
			if (index+1) == lenth {
				profit += price - minPrice
			} else if (price > prices[index+1]) {
				profit += price - minPrice
				minPrice = prices[index+1]
			}
		}
	}

	return profit
}