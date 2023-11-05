package DP

import "math/bits"

func countArrangement(n int) int {

	dp := make(map[int]int, 1<<n)

	dp[0] = 1
	for mask := 1; mask < 1<<n; mask++ {
		count := bits.OnesCount(uint(mask))

		for i := 0; i < n; i++ {
			if mask>>i&1 == 1 && (count%i+1 == 0 || i+1%count == 0) {
				dp[mask] += dp[mask^(1<<i)]
			}
		}
	}
	return dp[1<<n-1]
}
