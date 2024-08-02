package Hot100

func longestValidParentheses(s string) int {
	if s == "" {
		return 0
	}
	dp := make([]int, len(s)) //dp定义为以i为结尾的最长的有效括号

	for i := 1; i < len(s); i++ {
		if s[i] == ')' {
			if s[i-1] == '(' && i-2 >= 0 {
				dp[i] = dp[i-2] + 2
			} else {
				if i-1-dp[i-1] >= 0 && s[i-1-dp[i-1]] == '(' {
					if i-2-dp[i-1] >= 0 {
						dp[i] = dp[i-1] + 2 + dp[i-2-dp[i-1]]
					} else {
						dp[i] = dp[i-1] + 2
					}

				}
			}
		} else {
			continue
		}
	}

	res := -1
	for i := 0; i < len(dp); i++ {
		if dp[i] > res {
			res = dp[i]
		}
	}
	return res
}
