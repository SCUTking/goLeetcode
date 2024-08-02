package dynamic_programming

import "math"

func numSquares(n int) int {
	dp := make([]int, n+1) //dp[j]表示到达背包为j的最小的硬币数量

	coins := make([]int, 0)
	k := 1
	for k*k < n {
		coins = append(coins, k*k)
	}
	dp[0] = 0
	for i := 1; i < n+1; i++ {
		dp[i] = math.MaxInt32
	}
	for i := 0; i < len(coins); i++ {
		for j := coins[i]; j < n+1; j++ {
			if dp[j-coins[i]] != math.MaxInt32 {
				dp[j] = min(dp[j], dp[j-coins[i]]+1)
			}
		}
	}

	if dp[n] == math.MaxInt32 {
		return -1
	}
	return dp[n]
}
