package arr

import "math"

func divisionLand(land [][]int) int {
	m := len(land)
	n := len(land[0])

	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i > 0 {
				dp[i][j] += dp[i-1][j]
			}
			if j > 0 {
				dp[i][j] += dp[i][j-1]
			}
			if i > 0 && j > 0 {
				dp[i][j] -= dp[i-1][j-1]
			}
		}
	}
	res := math.MaxInt32
	for i := 0; i < m-1; i++ {
		diff := dp[m-1][n-1] - dp[m-1][i]
		if res > diff {
			res = diff
		}
	}
	for i := 0; i < n-1; i++ {
		diff := dp[m-1][n-1] - dp[i][n-1]
		if res > diff {
			res = diff
		}
	}
	return res
}
