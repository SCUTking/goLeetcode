package Monotone_stack

func trap(height []int) int {
	stack := make([]int, 0)
	//单调递减栈
	res := 0
	for i := 0; i < len(height); i++ {
		for len(stack) > 0 && height[i] > height[stack[len(stack)-1]] {
			i2 := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if len(stack) > 0 {
				i3 := stack[len(stack)-1]
				h := min(height[i3], height[i]) - height[i2]
				l := i - i3 - 1
				res += h * l
			}
		}
		stack = append(stack, i)
	}
	return res
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b

}
