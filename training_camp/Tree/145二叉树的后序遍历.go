package Tree

// 前序与后序的处理顺序和访问顺序是一致的。所以代码比较像
func postorderTraversal(root *TreeNode) []int {
	//中右左  反转就变成   左右中
	stack := make([]*TreeNode, 0)
	stack = append(stack, root)
	res := make([]int, 0)
	for len(stack) > 0 {
		head := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if head != nil {
			res = append(res, head.Val)
			stack = append(stack, head.Left)
			stack = append(stack, head.Right)
		}
	}
	for i := 0; i < len(res)/2; i++ {
		res[i], res[len(res)-i-1] = res[len(res)-i-1], res[i]
	}
	return res
}
