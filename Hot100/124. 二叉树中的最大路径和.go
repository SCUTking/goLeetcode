package Hot100

import "math"

func maxPathSum(root *TreeNode) int {
	var res int = math.MinInt32
	var dfs func(root *TreeNode) int
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		leftSum := max(dfs(root.Left), 0)
		rightSum := max(dfs(root.Right), 0)

		maxRes := rightSum + leftSum + root.Val
		if res < maxRes {
			res = maxRes
		}
		return root.Val + max(leftSum, rightSum)
	}
	dfs(root)
	return res
}
