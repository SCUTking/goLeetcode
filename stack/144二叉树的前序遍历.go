package stack

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) (res []int) {
	var stack []*TreeNode

	stack = append(stack, root)

	return nil
}
