package Hot100

func countBits(n int) []int {
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		if i&1 == 0 {
			//偶数
			dp[i] = dp[i>>1]
		} else {
			dp[i] = dp[i-1] + 1
		}
	}
	return dp

}
