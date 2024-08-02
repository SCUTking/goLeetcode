package dynamic_programming

import "fmt"

func change(amount int, coins []int) int {
	dp := make([]int, amount+1) //dp[j]表示到达背包为j的最小的硬币数量
	dp[0] = 1
	for i := 0; i < len(coins); i++ {
		for j := 0; j < amount+1; j++ {
			if j >= coins[i] {
				dp[j] = dp[j] + dp[j-coins[i]]
			}
		}
		fmt.Println(dp)
	}

	return dp[amount]

}
