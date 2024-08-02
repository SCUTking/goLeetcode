package dynamic_programming

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	dp[0][0] = 1
	for j := 0; j < m; j++ {
		for i := 0; i < n; i++ {
			if j-1 >= 0 && obstacleGrid[j-1][i] != 1 {
				dp[j][i] += dp[j-1][i]
			}
			if i-1 >= 0 && obstacleGrid[j-1][i] != 1 {
				dp[j][i] += dp[j][i-1]
			}
		}
	}
	return dp[m-1][n-1]
}
