package random2

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func convertBST(root *TreeNode) *TreeNode {

	//var getSum func(root *TreeNode,tar int)int
	nodes := make([]int, 0)
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Left)
		nodes = append(nodes, root.Val)
		dfs(root.Right)
	}

	var getSum func(target int) int
	getSum = func(target int) int {
		sum := 0
		for _, v := range nodes {
			if v >= target {
				sum += target
			}
		}
		return sum
	}

	var dfs1 func(root *TreeNode)
	dfs1 = func(root *TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Left)
		root.Val = getSum(root.Val)
		dfs(root.Right)
	}

	dfs(root)
	dfs1(root)

	return root

}
