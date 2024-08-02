package random1

func numOfSubarrays(nums []int) int {
	const modNum int = 1000000007
	odd, even := 0, 1
	ans := 0
	preSum := 0
	for i := 0; i < len(nums); i++ {
		preSum += nums[i]
		if preSum&1 == 1 {
			ans = (ans + even) % modNum
			odd++
		} else {
			ans = (ans + odd) % modNum
			even++
		}
	}

	return ans
}
