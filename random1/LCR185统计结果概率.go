package random1

import "fmt"

func statisticsProbability(num int) []float64 {

	dp := make([][]float64, num)
	for i := 0; i < num; i++ {

		dp[i] = make([]float64, 5*num+1)
	}
	for i := 0; i < 6; i++ {

		dp[0][i] = 1 / 6
		f := dp[0][i]
		fmt.Print(f)
	}

	for i := 1; i <= num; i++ {
		float64s := dp[i-1]
		for i2, f := range float64s {
			if f != 0 {
				for k := 0; k < 6; k++ {
					dp[i][i2+k] += f * 1 / 6
				}
			}
		}
	}

	return dp[num-1]

}
