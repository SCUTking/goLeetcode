package random2

func trap(height []int) int {
	//单调栈
	stack := make([]int, 0)
	ans := 0
	for i, h := range height {
		for len(stack) > 0 && h > height[stack[len(stack)-1]] {

			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if len(stack) <= 0 {
				break
			}
			left := stack[len(stack)-1]
			high := min(height[left], h) - top
			length := i - left - 1
			ans += high * length

		}
		stack = append(stack, i)
	}

	return ans

}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
