package tree

func IsSymmetric(root *TreeNode) bool {

	res := []int{}
	var dfs func(root *TreeNode)

	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Left)
		res = append(res, root.Val)
		dfs(root.Right)
	}

	dfs(root)

	if len(res)%2 == 0 {
		return false
	} else {
		for i := 0; i < len(res)/2; i++ {
			if res[i] != res[len(res)-i-1] {
				return false
			}
		}
		return true

	}

}
