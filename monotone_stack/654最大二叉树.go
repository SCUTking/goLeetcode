package monotone_stack

import (
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func constructMaximumBinaryTree(nums []int) *TreeNode {

	var dfs func(nums []int) *TreeNode
	dfs = func(nums []int) *TreeNode {
		if len(nums) == 0 {
			return nil
		}
		max := math.MinInt32
		in := 0
		for index, num := range nums {
			if num > max {
				max = num
				in = index
			}
		}

		root := &TreeNode{Val: max}
		root.Left = dfs(nums[:in])
		root.Right = dfs(nums[in+1:])

		return root

	}

	return dfs(nums)
}
