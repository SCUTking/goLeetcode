package Tree

func levelOrder(root *TreeNode) [][]int {
	queue := make([]*TreeNode, 0)
	if root == nil {
		return nil
	}
	queue = append(queue, root)
	res := make([][]int, 0)
	for len(queue) > 0 {
		temp := make([]*TreeNode, len(queue))
		copy(temp, queue)
		queue = queue[:0]
		tempArr := make([]int, 0)
		for _, node := range temp {
			tempArr = append(tempArr, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		res = append(res, tempArr)
	}

	return res
}
func getHigh1(node *TreeNode) int {
	if node == nil {
		return 0
	}
	return max1(getHigh1(node.Right), getHigh1(node.Left)) + 1
}

func max1(a, b int) int {
	if a > b {
		return a
	}
	return b
}
