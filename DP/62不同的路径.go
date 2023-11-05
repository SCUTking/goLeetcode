package DP

func uniquePaths(m int, n int) int {

	res := make([][]int, m)
	//初始化
	for i, _ := range res {
		var temp []int
		for i := 0; i < n; i++ {
			temp = append(temp, 1)
		}
		res[i] = temp
	}

	for j, v := range res {
		for i := 0; i < len(v); i++ {
			if j != 0 && i != 0 {
				v[i] = res[j-1][i] + res[j][i-1]
			}
		}
	}

	ints := res[m-1]
	return ints[n-1]

}
