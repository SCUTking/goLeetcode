package Hot100

func isMatch(s string, p string) bool {
	iLen := len(s)
	jLen := len(p)
	dp := make([][]bool, iLen) //s的前i个字符与p的前j个字符是否匹配
	for i := 0; i < iLen; i++ {
		dp[i] = make([]bool, jLen)
	}

	var match func(i, j int) bool
	match = func(i, j int) bool {
		if p[j] == '.' || p[j] == s[i] {
			return true
		}
		return false
	}

	dp[0][0] = true

	for i := 0; i < iLen; i++ {
		for j := 0; j < jLen; j++ {
			if p[j] == '.' {
				dp[i][j] = dp[i][j-1]
			} else if p[j] == '*' {
				if match(i, j-1) {
					if dp[i][j-2] || dp[i-1][j] {
						dp[i][j] = true
					}
				} else {
					dp[i][j] = dp[i][j-2]
				}
			} else {
				if match(i, j) {
					//
					dp[i][j] = dp[i-1][j-1]
				} else {
					dp[i][j] = false
				}
			}
		}
	}

	return dp[iLen-1][jLen-1]

}
