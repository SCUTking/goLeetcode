package DP

func Makesquare(matchsticks []int) bool {
	totalLen := 0
	for _, matchstick := range matchsticks {
		totalLen += matchstick
	}
	tlen := totalLen % 4

	if tlen != 0 {
		return false
	}
	dp := make([]int, 1<<len(matchsticks))
	//dp[0]=0
	for i := 1; i < len(dp); i++ {
		dp[1] = -1
	}

	for i := 1; i < len(dp); i++ {
		for k, v := range matchsticks {
			if i>>k&1 == 0 {
				continue
			}
			s1 := i &^ (1 << k)
			if dp[s1] >= 0 && dp[s1]+v <= tlen {
				dp[i] = (dp[s1] + v) % tlen
				break
			}

		}
	}

	return dp[len(dp)-1] == 0
}
