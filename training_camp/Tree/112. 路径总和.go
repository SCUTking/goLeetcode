package Tree

func hasPathSum(root *TreeNode, targetSum int) bool {
	var getPathSum func(pathSum int, node *TreeNode) bool
	getPathSum = func(pathSum int, node *TreeNode) bool {
		if node == nil {
			return false
		}
		pathSum += node.Val
		if node.Left == nil && node.Right == nil {
			if targetSum == pathSum {
				return true
			}
		}

		return getPathSum(pathSum, node.Right) || getPathSum(pathSum, node.Left)
	}

	return getPathSum(0, root)
}
