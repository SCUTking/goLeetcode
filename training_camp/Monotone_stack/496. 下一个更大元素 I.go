package Monotone_stack

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	m := make(map[int]int)

	stack := make([]int, 0)

	for i := 0; i < len(nums2); i++ {
		for len(stack) > 0 && stack[len(stack)-1] < nums2[i] {
			m[stack[len(stack)-1]] = nums2[i]
			stack = stack[:len(stack)-1]
		}

		stack = append(stack, nums2[i])
		if i == len(nums2)-1 {
			for j := 0; j < len(stack); j++ {
				m[stack[j]] = -1
			}
		}
	}

	res := make([]int, len(nums1))

	for i := 0; i < len(nums1); i++ {
		res[i] = m[nums1[i]]
	}
	return res
}
