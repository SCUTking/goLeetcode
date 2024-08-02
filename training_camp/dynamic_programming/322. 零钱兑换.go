package dynamic_programming

import "math"

// 完全背包问题
func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	dp := make([]int, amount+1) //dp[j]表示到达背包为j的最小的硬币数量
	dp[0] = 0
	for i := 0; i < len(coins); i++ {
		for j := 0; j < amount+1; j++ {
			if j >= coins[i] {
				if dp[j] == 0 {
					if dp[j-coins[i]] != 0 || j-coins[i] == 0 {
						dp[j] = dp[j-coins[i]] + 1
					}

				} else {
					if dp[j-coins[i]] != 0 || j-coins[i] == 0 {
						dp[j] = min(dp[j], dp[j-coins[i]]+1)
					}
				}
			}
		}
	}

	if dp[amount] == 0 {
		return -1
	}
	return dp[amount]
}

// 完全背包问题
func coinChange1(coins []int, amount int) int {
	dp := make([]int, amount+1) //dp[j]表示到达背包为j的最小的硬币数量
	dp[0] = 0
	for i := 1; i < amount+1; i++ {
		dp[i] = math.MaxInt32
	}
	for i := 0; i < len(coins); i++ {
		for j := coins[i]; j < amount+1; j++ {
			if dp[j-coins[i]] != math.MaxInt32 {
				dp[j] = min(dp[j], dp[j-coins[i]]+1)
			}
		}
	}

	if dp[amount] == math.MaxInt32 {
		return -1
	}
	return dp[amount]
}
