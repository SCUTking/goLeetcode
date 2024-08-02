package dynamic_programming

import "math"

func maxProfit(prices []int) int {
	minNum := math.MaxInt32
	res := 0
	for i := 0; i < len(prices); i++ {
		if prices[i] < minNum {
			minNum = prices[i]
		}
		res = max(res, prices[i]-minNum)
	}
	return res
}
