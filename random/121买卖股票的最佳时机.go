package random

import (
	"math"
	"sort"
)

//func maxProfit(prices []int) int {
//	ans := 0
//	for i, price := range prices {
//		for j := i; j < len(prices); j++ {
//			if price < prices[j] {
//				ans = max3(ans, prices[j]-price)
//			}
//		}
//	}
//	return ans
//}

func MaxProfit(prices []int) int {
	ans := 0
	min := math.MaxInt32
	for i, price := range prices {
		if price >= min {
			continue
		}
		temp := make([]int, len(prices)-1-i)
		if i == len(prices)-1 {
			break
		}
		copy(temp, prices[i+1:])
		sort.Ints(temp)
		max := temp[len(prices)-1]
		if price < max {
			ans = max3(ans, max-price)
			min = price
		}

	}
	return ans
}

func max3(a, b int) int {
	if a > b {
		return a
	}
	return b
}
