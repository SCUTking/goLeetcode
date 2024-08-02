package Hot100

import "math"

func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	dp[0][0] = 0
	for i := 0; i < m; i++ {
		dp[i][0] = i
	}
	for i := 0; i < n; i++ {
		dp[0][i] = i
	}

	//操作转化为3个，在A后面插入一个字符，在B后面插入，修改A，B后面的

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if word1[i] == word2[j] {
				dp[i][j] = min1(dp[i-1][j], dp[i][j-1], dp[i-1][j-1])
			} else {
				dp[i][j] = min1(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]) + 1
			}

		}
	}

	return dp[m-1][n-1]
}

func min1(args ...int) int {
	temp := math.MaxInt32
	for i := 0; i < len(args); i++ {
		if args[i] < temp {
			temp = args[i]
		}
	}

	return temp
}
