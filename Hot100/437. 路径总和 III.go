package Hot100

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) int {
	res := 0
	preSum := make(map[int64]int, 0) //k是前缀和  value是该前缀和的数量
	preSum[0] = 1                    //要记录空路径的数量为1
	var dfs func(root *TreeNode, pre int64)
	dfs = func(root *TreeNode, pre int64) {
		if root == nil {
			return
		}
		pre += int64(root.Val)
		res += preSum[pre-int64(targetSum)]

		preSum[pre]++
		dfs(root.Left, pre)
		dfs(root.Right, pre)
		preSum[pre]--
	}
	dfs(root, 0)
	return res
}
func myPathSum(root *TreeNode, targetSum int) int {
	res := 0
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		res += getSumOfRoot(root, targetSum)
		dfs(root.Left)
		dfs(root.Right)
	}
	return res
}

func getSumOfRoot(root *TreeNode, targetSum int) int {
	res := 0
	var dfs func(root *TreeNode, pre int)
	dfs = func(root *TreeNode, pre int) {
		if root == nil {
			return
		}
		pre += root.Val
		if pre == targetSum {
			res++
		}
		dfs(root.Left, pre)
		dfs(root.Right, pre)
	}

	dfs(root, 0)
	return res
}
