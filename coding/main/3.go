package main

import (
	"fmt"
)

func main() {
	n := 0

	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}

	fmt.Scan(&n)
	if n < 2 {
		return
	}
	x := 0
	for x < n {
		//var s string
		//fmt.Scan(&s)

		for j := 0; j < n; j++ {
			var num int
			fmt.Scan(&num)

			if x > 1 && j > 1 {
				dp[x][j] = dp[x-1][j] + dp[x][j-1] - dp[x-1][j-1]
			} else if x == 0 && j > 1 {
				dp[x][j] = dp[x][j-1]
			} else if j == 0 && x > 1 {
				dp[x][j] = dp[x-1][j]
			} else {

			}
			if num == 0 {
				dp[x][j]++
			}
		}
		x++
	}
	for i := 1; i <= n; i++ {
		if i&1 == 0 {
			count := 0
			for j := i - 1; j < n; j++ {
				for k := i - 1; k < n; k++ {
					var num0 int
					if k >= i {
						num0 = dp[j][k] - dp[j][k-i]
					} else {
						num0 = dp[j][k]
					}

					if i*i-num0 == num0 {
						count++
					}
				}
			}
			fmt.Println(count)
		} else {
			fmt.Println(0)
		}

	}
}
