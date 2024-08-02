package dynamic_programming

func maxProfit1(prices []int) int {
	//dp[i][0] 表示的持有股票的最大现金
	//i-1 也持有
	//i 刚买入

	//dp[i][1] 表示不持有股票的最大现金
	//i-1也不持有
	//i 刚卖掉

	dp := make([][2]int, len(prices))
	//初始化
	dp[0][0] = -prices[0]
	dp[0][1] = 0
	for i := 1; i < len(prices); i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]-prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]+prices[i])
	}

	return dp[len(prices)-1][1]
}
