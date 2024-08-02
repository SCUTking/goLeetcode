package dynamic_programming

import "strings"

func findMaxForm(strs []string, m int, n int) int {
	//01背包问题
	//背包的大小m个0 n个1
	//物品的重量
	//01只能取或不取
	l := len(strs)
	dp := make([][][]int, n)
	for i := 0; i < l; i++ {
		dp[i] = make([][]int, m+1)
		for j := 0; j < m+1; j++ {
			dp[i][j] = make([]int, n+1)
		}
	}

	nums := make([]ZO, l)
	for i, str := range strs {
		count1 := strings.Count(str, "1")
		count0 := len(str) - count1
		nums[i] = ZO{one: count1, zero: count0}
	}

	for i := nums[0].zero; i < m+1; i++ {
		for j := nums[0].one; j < n+1; j++ {
			dp[0][i][j] = 1
		}
	}

	for i := 1; i < l; i++ {
		for j := 0; j < m+1; j++ {
			for k := 0; k < n+1; k++ {
				if j < nums[i].zero || k < nums[i].one {
					dp[i][j][k] = dp[i-1][j][k]
				} else {
					dp[i][j][k] = max(dp[i-1][j][k], dp[i-1][j-nums[i].zero][k-nums[i].one]+1)
				}
			}
		}
	}

	return dp[l-1][m][n]
}

type ZO struct {
	zero int
	one  int
}
