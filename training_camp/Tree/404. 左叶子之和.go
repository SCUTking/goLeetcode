package Tree

func sumOfLeftLeaves(root *TreeNode) int {
	var res int
	var getLeft func(node *TreeNode)
	getLeft = func(node *TreeNode) {
		if node == nil {
			return
		}

		if node.Left != nil && node.Left.Left == nil && node.Left.Right == nil {
			res += node.Left.Val
		}
		getLeft(node.Left)
		getLeft(node.Right)
	}
	getLeft(root)
	return res
}
