package DP

func IntegerReplacement(n int) int {
	//dp := make([]int, n+1)
	//dp[0] = 0
	//dp[1] = 0
	//
	//for i := 2; i < n+1; i++ {
	//	if i&1 == 0 {
	//		//说明是偶数
	//		dp[i] = dp[i/2] + 1
	//	} else {
	//		dp[i] = dp[i-1] + 1
	//	}
	//}
	//for i, v := range dp {
	//	fmt.Println(i)
	//	print("->")
	//	print(v)
	//}
	//return dp[n]

	m := make(map[int]int)
	var helper func(n int) int
	helper = func(n int) (res int) {
		if n == 1 {
			return 0
		}
		if res, ok := m[n]; ok {
			return res
		}
		defer func() {
			m[n] = res
		}()
		if n&1 == 0 {
			res = helper(n/2) + 1
		} else {
			//return 2 + min(helper(n/2), helper(n/2+1))
			res = 1 + min(helper(n+1), helper(n-1))
		}
		return
	}

	return helper(n)

}

func min(int2 int, int3 int) int {
	if int3 > int2 {
		return int2
	}
	return int3
}
