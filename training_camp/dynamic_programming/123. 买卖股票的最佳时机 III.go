package dynamic_programming

func maxProfit2(prices []int) int {

	n := len(prices)
	dp := make([][5]int, n)
	// 0 没有操作
	// 1 持有第一次的股票
	// 2 卖掉第一次的股票
	// 3 持有第二次的股票
	// 4 卖掉第二次的股票
	dp[0][0] = 0
	dp[0][1] = -prices[0]
	dp[0][2] = 0
	dp[0][3] = -prices[0]
	dp[0][4] = 0

	for i := 1; i < n; i++ {
		dp[i][0] = dp[i-1][0]
		dp[i][1] = max(dp[i-1][0]-prices[i], dp[i-1][1])
		dp[i][2] = max(dp[i-1][1]+prices[i], dp[i-1][2])
		dp[i][3] = max(dp[i-1][2]-prices[i], dp[i-1][3])
		dp[i][4] = max(dp[i-1][3]+prices[i], dp[i-1][4])
	}

	return dp[n-1][4]
}
