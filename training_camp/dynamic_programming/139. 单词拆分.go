package dynamic_programming

import "fmt"

func wordBreak(s string, wordDict []string) bool {

	dp := make([]bool, len(s)+1)
	dp[0] = true
	m := make(map[string]struct{})
	for i := 0; i < len(wordDict); i++ {
		m[wordDict[i]] = struct{}{}
	}

	for j := 0; j < len(s)+1; j++ {
		for i := 0; i < len(wordDict); i++ {
			_, ok := m[s[j-len(wordDict[i]):j]]
			if j >= len(wordDict[i]) {
				dp[j] = dp[j] || (dp[j-len(wordDict[i])] && ok)
			}
		}
		fmt.Println(dp)
	}

	return dp[len(s)]
}
