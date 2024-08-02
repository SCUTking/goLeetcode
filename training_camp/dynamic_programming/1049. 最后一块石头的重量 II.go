package dynamic_programming

func lastStoneWeightII(stones []int) int {
	//01背包问题
	//定义石头总重为sum 考虑将这堆石头分为两堆 当两堆的之间的差最小时 就可以得出最后一个石头的重量
	//用dp[i][j]前i个石头，是否能凑重量为j的

	sum := 0
	for _, stone := range stones {
		sum += stone
	}
	n := len(stones)
	dp := make([][]bool, n+1)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, sum/2+1)
	}

	dp[0][0] = true
	for i := 0; i < n; i++ {
		for j := 0; j < sum/2+1; j++ {
			if j < stones[i] {
				dp[i+1][j] = dp[i][j]
			} else {
				dp[i+1][j] = dp[i][j-stones[i]] || dp[i][j]
			}
		}
	}

	for i := sum / 2; i >= 0; i-- {
		if dp[n][i] {
			return sum - 2*i
		}
	}

	return 0

}
