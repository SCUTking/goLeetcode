package coding

//3. 最长回文子串：
//给你一个字符串s，找到 s 中最长的回文子串。
//示例：输入：s = "babad" 输出："bab" 解释："aba" 同样是符合题意的答案。
//输入：s = "cbbd" 输出："bb"

func findLongest(s string) (res string) {
	n := len(s)
	dp := make([][]bool, n)

	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
		dp[i][i] = true
	}
	if n <= 1 {
		return s
	}
	dp[0][0] = true
	if s[0] == s[1] {
		dp[0][1] = true
	}

	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			if i == j {
				dp[i][j] = true
			} else {
				if j-1 > 0 {
					dp[i][j] = dp[i+1][j-1] && s[i] == s[j]
				}
			}

			if dp[i][j] {
				if j-i+1 > len(res) {
					res = s[i : j+1]
				}
			}

		}
	}

	return res

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
