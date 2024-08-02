package dynamic_programming

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rob3(root *TreeNode) int {

	//返回值第一个参数选择的最大值，第二个是不选择的最大值
	var dfs func(root *TreeNode) (int, int)
	dfs = func(root *TreeNode) (int, int) {

		if root.Left == nil && root.Right == nil {
			return root.Val, 0
		}
		noSum := 0  //不选择这一层
		maxSum := 0 //选择这一层
		if root.Left != nil {
			i2, no := dfs(root.Left)
			maxSum += i2
			noSum += no
		}
		if root.Right != nil {
			i1, no := dfs(root.Right)
			maxSum += i1
			noSum += no
		}

		return max(noSum+root.Val, maxSum), maxSum
	}

	res, _ := dfs(root)
	return res
}
func rob5(root *TreeNode) int {

	//返回值第一个参数选择的最大值，第二个是不选择的最大值
	var dfs func(root *TreeNode) (int, int)
	dfs = func(root *TreeNode) (int, int) {
		if root == nil {
			return 0, 0
		}

		i, i2 := dfs(root.Left)
		i3, i4 := dfs(root.Right)

		//选择
		res1 := root.Val + i2 + i4
		//不选择  那么子节点也是可以选择或者不选的  就看最大值
		res2 := max(i, i2) + max(i3, i4)

		return res1, res2
	}

	res1, res2 := dfs(root)
	return max(res1, res2)
}
