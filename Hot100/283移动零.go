package Hot100

func moveZeroes(nums []int) {

	count := 0
	for _, num := range nums {
		if num == 0 {
			count++
		}
	}

	for i := 0; i < count; i++ {
		if nums[i] == 0 {
			for j := i; j < len(nums); j++ {
				if nums[j] != 0 {
					nums[j], nums[i] = nums[i], nums[j]
					break
				}
			}
		}
	}

	return
}

func moveZeroes1(nums []int) {

	left, right, n := 0, 0, len(nums)
	for right < n {
		if nums[right] != 0 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
		right++
	}

}
