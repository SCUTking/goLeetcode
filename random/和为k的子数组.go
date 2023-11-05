package random

func SubarraySum(nums []int, k int) int {
	//curSum := 0
	//ans1 := 0
	//left, right := 0, 0
	//for right < len(nums) {
	//	curSum += nums[right]
	//	right++
	//	for left < right && curSum >= k {
	//		if curSum == k {
	//			ans1++
	//		}
	//		curSum -= nums[left]
	//		left++
	//	}
	//
	//}
	//
	//return ans1
	ans := 0
	n := len(nums)
	for i := 1; i <= n; i++ {
		left, right := 0, 0
		cur := 0
		for right < n {
			cur += nums[right]
			right++
			for right-left > i {
				cur -= nums[left]
				left++
			}
			if cur == k && right-left == i {
				ans++
			}
		}
	}

	return ans
}
