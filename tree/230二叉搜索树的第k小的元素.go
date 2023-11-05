package tree

func kthSmallest(root *TreeNode, k int) int {

	ans := 0
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Left)
		k--
		if k == 0 {
			ans = root.Val
		}
		dfs(root.Right)
	}

	dfs(root)
	return ans
}
