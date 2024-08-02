package Tree

func inorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Left)
		res = append(res, root.Val)
		dfs(root.Right)
	}
	dfs(root)
	return res

}

// 迭代版本
func inorderTraversal1(root *TreeNode) []int {
	//中左右
	stack := make([]*TreeNode, 0)
	cur := root
	res := make([]int, 0)
	for len(stack) > 0 || cur != nil {
		if cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		} else {
			head := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res = append(res, head.Val)
			cur = head.Right
		}
	}
	return res
}
