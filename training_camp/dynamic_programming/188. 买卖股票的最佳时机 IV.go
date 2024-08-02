package dynamic_programming

func maxProfit3(k int, prices []int) int {
	n := len(prices)
	dp := make([][]int, n)
	l := 2*k + 1
	for i := 0; i < n; i++ {
		dp[i] = make([]int, l)
	}

	for i := 1; i < l; i += 2 {
		dp[0][i] = -prices[0]
	}

	for i := 1; i < n; i++ {
		for j := 1; j < l; j++ {
			if j&1 == 0 {
				//这个操作是买入
				dp[i][j] = max(dp[i-1][j-1]-prices[i], dp[i-1][j])
			} else {
				dp[i][j] = max(dp[i-1][j-1]+prices[i], dp[i-1][j])
			}

		}
	}

	return dp[n-1][l-1]
}
