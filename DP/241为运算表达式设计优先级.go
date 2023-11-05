package DP

import (
	"unicode"
)

const addition, subtraction, multiplication = -1, -2, -3

func diffWaysToCompute(expression string) []int {

	opt := []int{}

	for i, n := 0, len(expression); i < n; {

		if unicode.IsDigit(rune(expression[i])) {
			x := 0
			for ; unicode.IsDigit(rune(expression[i])) && i < n; i++ {
				x = x*10 + int(expression[i]-'0')
			}
			opt = append(opt, x)

		} else {
			if expression[i] == '+' {
				opt = append(opt, addition)
			} else if expression[i] == '-' {
				opt = append(opt, subtraction)
			} else {
				opt = append(opt, multiplication)
			}
			i++

		}
	}

	n := len(opt)
	dp := make([][][]int, n)
	for i, _ := range dp {
		dp[i] = make([][]int, n)
	}

	var dfs func(int, int) []int

	dfs = func(l int, r int) []int {
		res := dp[l][r]
		if res != nil {
			return res
		}
		if l == r {
			//此时只有有一个数字，没有操作符
			dp[l][r] = []int{opt[l]}
			return dp[l][r]
		}

		for i := l; i < r-1; i += 2 {
			left := dfs(l, i)
			right := dfs(i+2, r)
			for _, x := range left {
				for _, y := range right {
					if opt[i+1] == addition {
						dp[l][r] = append(dp[l][r], x+y)
					} else if opt[i+1] == subtraction {
						dp[l][r] = append(dp[l][r], x-y)
					} else {
						dp[l][r] = append(dp[l][r], x*y)
					}
				}
			}

		}
		return dp[l][r]
	}

	return dfs(0, n-1)

}
