package tree

func pathSum(root *TreeNode, targetSum int) (res [][]int) {
	path := []int{}
	var dfs func(root *TreeNode, left int)
	dfs = func(root *TreeNode, left int) {
		if root == nil {
			return
		}
		path = append(path, root.Val)
		left -= root.Val
		defer func() {
			path = path[:len(path)-1]
		}()
		if root.Left == nil && root.Right == nil && left == 0 {
			res = append(res, append([]int(nil), path...))
			return
		}
		dfs(root.Left, left)
		dfs(root.Right, left)

	}

	dfs(root, targetSum)

	return

}

func find(root *TreeNode, add int, path *[]int, targetSum int, res *[][]int) {

	if root == nil {
		return
	}
	ints := *path
	ints = append(ints, root.Val)

	if root.Left != nil {
		find(root.Left, add+root.Val, &ints, targetSum, res)
	}
	if root.Left == nil && root.Right == nil {
		if add+root.Val == targetSum {
			r := *res
			r = append(r, ints)
			ints = []int{}

		}
	}
	if root.Right != nil {

		find(root.Left, add+root.Val, &ints, targetSum, res)
	}

}
