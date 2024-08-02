package dynamic_programming

func uniquePaths(m int, n int) int {
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	dp[0][0] = 1
	for j := 0; j < m; j++ {
		for i := 0; i < n; i++ {
			if j-1 >= 0 {
				dp[j][i] += dp[j-1][i]
			}
			if i-1 >= 0 {
				dp[j][i] += dp[j][i-1]
			}
		}
	}
	return dp[m-1][n-1]
}
