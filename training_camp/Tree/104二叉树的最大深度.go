package Tree

func maxDepth(root *TreeNode) int {
	return getHigh(root)
}

func getHigh(node *TreeNode) int {
	if node == nil {
		return 0
	}
	return Max(getHigh(node.Left), getHigh(node.Right)) + 1
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
