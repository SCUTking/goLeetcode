package Tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		res = append(res, root.Val)
		dfs(root.Left)
		dfs(root.Right)
	}
	dfs(root)
	return res
}

func preorderTraversal1(root *TreeNode) []int {
	//中左右
	stack := make([]*TreeNode, 0)
	stack = append(stack, root)
	res := make([]int, 0)
	for len(stack) > 0 {
		head := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if head != nil {
			res = append(res, head.Val)
			stack = append(stack, head.Right)
			stack = append(stack, head.Left)
		}
	}

	return res
}
