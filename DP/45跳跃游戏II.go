package DP

func Jump(nums []int) int {

	num0 := []int{0}
	if len(nums) == 1 {
		return 0
	}
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			var length int = i - j
			if nums[j] >= length {
				num0 = append(num0, num0[j]+1)
				break
			}
		}
	}

	return 1

}
