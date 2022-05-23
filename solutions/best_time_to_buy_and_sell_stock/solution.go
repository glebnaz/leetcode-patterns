package best_time_to_buy_and_sell_stock

func maxProfit(prices []int) int {
	minPrice := prices[0]

	maxProfit := 0

	for _, v := range prices {
		if maxProfit < v-minPrice {
			maxProfit = v - minPrice
		}

		if minPrice > v {
			minPrice = v
		}
	}

	return maxProfit
}
