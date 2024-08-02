package Tree

func isBalanced(root *TreeNode) bool {
	return getNodeHigh(root) >= 0
}
func getNodeHigh(node *TreeNode) int {
	if node == nil {
		return 0
	}
	left := getNodeHigh(node.Left)
	right := getNodeHigh(node.Right)
	if left == -1 || right == -1 || abs(left-right) > 1 {
		return -1
	}

	return myMax(left, right) + 1
}

func myMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
