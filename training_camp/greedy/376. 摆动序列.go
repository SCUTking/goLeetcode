package greedy

func wiggleMaxLength(nums []int) int {
	temp := make([]int, 0)
	for i, _ := range nums {
		if i > 0 {
			temp = append(temp, nums[i]-nums[i-1])
		}
	}
	if len(temp) == 0 {
		return 1
	}

	res := 0
	pre := 0
	for i := 0; i < len(temp); i++ {
		if temp[i] > 0 {
			if pre == 1 {
				continue
			}
			pre = 1
			res++
		} else if temp[i] < 0 {
			if pre == 0 {
				continue
			}
			pre = 0
			res++
		} else {
			continue
		}
	}
	return res + 1
}
