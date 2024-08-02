package Monotone_stack

func largestRectangleArea(heights []int) int {

	s := make([]int, 1)
	s = append(s, heights...)
	s = append(s, 0)

	res := 0
	stack := make([]int, 0) //单调递增栈  当小于栈顶元素时触发
	for i := 0; i < len(s); i++ {
		for len(stack) > 0 && s[stack[len(stack)-1]] > s[i] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			if len(stack) > 0 {
				left := stack[len(stack)-1]

				h := s[top]
				w := i - left
				res = max(h*w, res)
			}
		}
		stack = append(stack, i)
	}

	return res

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
