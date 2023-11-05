package tree

func ZigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	queue := []*TreeNode{root}
	res := make([][]int, 1)

	var helper func(level int)
	helper = func(level int) {
		n := len(queue)
		if n == 0 {
			return
		}
		if len(res) < level+1 {
			res = append(res, []int{})
		}

		if level%2 != 0 {
			for i := len(queue) - 1; i >= 0; i-- {
				res[level] = append(res[level], queue[i].Val)
			}
		}

		for i := 0; i < n; i++ {
			cur := queue[0]
			queue = queue[1:]
			if level%2 == 0 {
				res[level] = append(res[level], cur.Val)
			}

			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}
			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}
		}

		helper(level + 1)

	}

	helper(0)
	return res
}
