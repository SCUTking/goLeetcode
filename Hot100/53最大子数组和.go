package Hot100

import "math"

func maxSubArray(nums []int) int {
	n := len(nums)
	preSum := make([]int, n+1)
	minSum := math.MinInt32
	res := math.MinInt32
	for i := 1; i <= n; i++ {
		preSum[i] = preSum[i-1] + nums[i-1]
		sum := preSum[i] - minSum
		if sum > res {
			res = sum
		}
		if minSum > preSum[i] {
			minSum = preSum[i]
		}

	}
	return res

}
