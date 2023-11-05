package random

func integerBreak(n int) int {
	dp := make(map[int]int)
	dp[0] = 0
	dp[1] = 0
	dp[2] = 1
	if n <= 2 {
		return dp[n]
	}
	for i := 3; i <= n; i++ {
		cur := 0
		for j := 1; j < i; j++ {
			i2 := j * (i - j)
			i3 := j * dp[i-j]
			i4 := max1(i3, i2)
			cur = max1(i4, cur)
		}
		dp[i] = cur % 1000000007
	}

	return dp[n]

}

func max1(a, b int) int {
	if a < b {
		return b
	}
	return a
}
