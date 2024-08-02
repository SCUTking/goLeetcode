package Tree

// 翻转二叉树，翻转节点其左右子树节点也会被翻转过来的。
// 所以直接递归翻转就好

func invertTree(root *TreeNode) *TreeNode {
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		node.Right, node.Left = node.Left, node.Right
		dfs(node.Left)
		dfs(node.Right)
	}
	dfs(root)
	return root
}
