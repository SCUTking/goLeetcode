package Hot100

func diameterOfBinaryTree(root *TreeNode) int {
	res := 0
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		cur := getMaxPath(root.Left) + getMaxPath(root.Right)
		if cur > res {
			res = cur
		}
		dfs(root.Left)
		dfs(root.Right)
	}
	dfs(root)
	return res
}

func getMaxPath(root *TreeNode) int {
	var dfs func(root *TreeNode, pre int) int
	dfs = func(root *TreeNode, pre int) int {
		if root == nil {
			return pre
		}
		pre++
		return max(dfs(root.Left, pre), dfs(root.Right, pre))
	}
	return dfs(root, 0)
}

var ans int

func getDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	L := getDepth(root.Left)
	R := getDepth(root.Right)
	ans = max(ans, L+R+1)
	return max(L, R) + 1
}
