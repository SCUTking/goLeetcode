package random2

func subarraySum(nums []int, k int) int {

	count := 0

	for i := 1; i <= len(nums); i++ {
		left, right := 0, 0
		sum := 0
		for right < len(nums) {
			sum += nums[right]
			right++
			for right-left >= i {
				if sum == k {
					count++
				}
				sum -= nums[left]
				left++
			}

		}
	}

	return count
}
