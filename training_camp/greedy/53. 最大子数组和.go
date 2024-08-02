package greedy

import (
	"math"
)

func maxSubArray(nums []int) int {
	n := len(nums)
	preSum := make([]int, n)
	preSum[0] = nums[0]
	for i := 1; i < n; i++ {
		preSum[i] = preSum[i-1] + nums[i]
	}
	minPreSum := 0
	res := math.MinInt32
	for i := 0; i < n; i++ {
		temp := preSum[i] - minPreSum
		minPreSum = min(preSum[i], minPreSum)
		res = max(temp, res)
	}

	return res
}
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
