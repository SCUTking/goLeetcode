package prefixAnd

func productExceptSelf(nums []int) []int {
	prefix := make([]int, 0)
	after := make([]int, 0)
	prefix = append(prefix, 1)
	after = append(after, 1)
	//前缀
	temp := 1
	for i := 0; i < len(nums); i++ {
		temp *= nums[i]
		prefix = append(prefix, temp)
	}
	//后缀
	temp2 := 1
	for i := len(nums) - 1; i >= 0; i-- {
		temp2 *= nums[i]
		after[i] = temp2
	}
	prefix = append(prefix, 1)
	after = append(after, 1)

	res := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		res = append(res, prefix[0]*after[i+1])
	}
	return res

}
