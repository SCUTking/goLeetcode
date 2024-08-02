package Hot100

func maxProfit(prices []int) int {
	n := len(prices)
	dp := make([][3]int, n) //第i天手上的最大的钱
	dp[0][0] = -prices[0]   //持有
	dp[0][1] = 0            //不持有,不处于冷冻期，明天能买
	dp[0][2] = 0            //不持有，处于冷冻期,明天不能买
	for i := 1; i < n; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]-prices[i])
		dp[i][1] = max(dp[i-1][2], dp[i-1][1])
		dp[i][2] = dp[i-1][0] + prices[i]
	}

	return max(dp[n-1][1], dp[n-1][2])
}
