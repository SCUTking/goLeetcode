package Hot100

import "math"

func maximalSquare(matrix [][]byte) int {
	m := len(matrix)
	n := len(matrix[0])
	dp := make([][]int, m) //以[i][j]为右下角，所能构建最大的正方形的长度
	//初始化
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	res := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == '1' {
				if i == 0 || j == 0 {
					dp[i][j] = 1
				} else {
					dp[i][j] = minOf(dp[i-1][j-1], dp[i-1][j], dp[i][j-1]) + 1
				}
			} else {
				dp[i][j] = 0
			}
			if res < dp[i][j] {
				res = dp[i][j]
			}
		}
	}
	return res * res

}

func minOf(args ...int) int {
	res := math.MaxInt32
	for i := 0; i < len(args); i++ {
		if res > args[i] {
			res = args[i]
		}
	}
	return res
}
