package greedy

// 动态规划
func maxProfit(prices []int) int {
	n := len(prices)
	dp := make([][2]int, n)
	dp[0][0] = 0          //表示没有股票
	dp[0][1] = -prices[0] //表示持有股票
	for i := 1; i < n; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
	}
	return dp[n-1][0]
}
