package Tree

func convertBST(root *TreeNode) *TreeNode {
	var dfs func(node *TreeNode)
	sum := 0
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Right)
		sum += node.Val
		node.Val = sum
		dfs(node.Left)
	}
	dfs(root)
	return root
}
