package random

import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func WidthOfBinaryTree(root *TreeNode) int {
	type pair struct {
		index int
		cur   *TreeNode
	}

	queue := []pair{{1, root}}
	ans := math.MinInt32
	for len(queue) > 0 {
		ans = max2(ans, queue[len(queue)-1].index-queue[0].index+1)
		tmp := queue
		queue = queue[:0]
		for _, p := range tmp {
			if p.cur.Left != nil {
				queue = append(queue, pair{index: p.index * 2, cur: p.cur.Left})
			}
			if p.cur.Right != nil {
				queue = append(queue, pair{index: p.index*2 + 1, cur: p.cur.Right})
			}

		}

	}

	return ans

}
func max2(a, b int) int {
	if a > b {
		return a
	}
	return b

}
