package Hot100

func countSubstrings(s string) int {
	n := len(s)
	dp := make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
	}
	res := 0
	for i := 0; i < n; i++ {
		dp[i][i] = true
	}
	res += n

	for i := 0; i < n-1; i++ {
		if s[i] == s[i+1] {
			dp[i][i+1] = true
			res++
		}
	}

	if n < 3 {
		return res
	}
	for i := 3; i < n; i++ {
		for j := 0; j < n-2; j++ {
			if s[j] == s[j+i-1] && dp[j+1][j+i-2] {
				dp[j][j+i-1] = true
				res++
			}
		}
	}

	return res

}
