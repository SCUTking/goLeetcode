package greedy

import "math"

func maxProfit1(prices []int) int {
	res := math.MinInt32
	minIn := prices[0]
	for _, price := range prices {
		temp := price - minIn
		res = max(temp, res)
		minIn = min(minIn, price)
	}
	return res
}
