package Tree

//一开始看题目表示无虫下手，想了十分钟也是，就看题解了。发现还是不会拆解子问题。

// 递归版本————
// 转化为：两个树什么时候互为镜像
// 根节点相同
// 一个数的左指数等于另一颗树的右指数
func isSymmetric(root *TreeNode) bool {
	return Check(root, root)
}

func Check(node1 *TreeNode, node2 *TreeNode) bool {

	//判断是否为空
	if node1 == nil && node2 == nil {
		return true
	}

	if node1 == nil || node2 == nil {
		return false
	}

	return node1.Val == node2.Val && Check(node1.Left, node2.Right) && Check(node1.Right, node2.Left)

}

// 迭代——当队列为空时，或者我们检测到树不对称（即从队列中取出两个不相等的连续结点）时，该算法结束
func isSymmetric1(root *TreeNode) bool {
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	queue = append(queue, root)

	for len(queue) > 0 {
		node1 := queue[0]
		node2 := queue[1]
		queue = queue[2:]
		if node1 == nil && node2 == nil {
			continue
		}

		if node1 == nil || node2 == nil {
			return false
		}

		if node1.Val != node2.Val {
			continue
		}

		//将两指针的左右 节点 作为一对同时处理
		queue = append(queue, node1.Right)
		queue = append(queue, node2.Left)

		queue = append(queue, node1.Left)
		queue = append(queue, node2.Right)

	}

	return true
}
