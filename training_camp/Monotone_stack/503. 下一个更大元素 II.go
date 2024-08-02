package Monotone_stack

func nextGreaterElements(nums []int) []int {
	nums2 := make([]int, len(nums))

	copy(nums2, nums)

	for i := 0; i < len(nums); i++ {
		nums2 = append(nums2, nums[i])
	}

	return nextGreaterElement(nums, nums2)
}
func nextGreaterElements1(nums []int) []int {
	n := len(nums)
	res := make([]int, n)
	//初始化
	for i := 0; i < n; i++ {
		res[i] = -1
	}
	stack := make([]int, 0)
	for i := 0; i < n*2; i++ {
		for len(stack) > 0 && nums[stack[len(stack)-1]] < nums[i%n] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res[top] = nums[i%n]
		}
		stack = append(stack, i%n)
	}

	return res
}
