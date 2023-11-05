package DP

func WordBreak(s string, wordDict []string) bool {
	m := make(map[string]bool)
	for _, s := range wordDict {
		m[s] = true
	}

	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if dp[j] && m[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}

	return dp[len(s)]

}
