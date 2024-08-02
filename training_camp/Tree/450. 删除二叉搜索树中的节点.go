package Tree

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	cur := root
	var parent *TreeNode
	for cur != nil && cur.Val != key {
		if cur.Val < key {
			parent = cur
			cur = cur.Right
		} else {
			parent = cur
			cur = cur.Left
		}
	}
	if cur == nil {
		return root
	}
	var temp *TreeNode
	if cur.Left == nil && cur.Right == nil {
		temp = nil
	}

	if cur.Left != nil && cur.Right == nil {
		temp = cur.Left
	} else if cur.Right != nil && cur.Left == nil {
		temp = cur.Right
	}

	if parent == nil {
		return temp
	}

	if cur.Right != nil && cur.Left != nil {
		var minRight *TreeNode
		minRight = cur.Right
		for minRight.Left != nil {
			minRight = minRight.Left
		}
		cur.Val = minRight.Val
		node := deleteNode(cur.Right, minRight.Val)
		cur.Right = node
	} else {
		if parent.Left == cur {
			parent.Left = temp
		} else {
			parent.Right = temp
		}
	}
	return root
}

// 递归的实现 删除节点
func deleteNode1(root *TreeNode, key int) *TreeNode {
	switch {
	case root == nil:
		return nil
	case root.Val > key:
		root.Left = deleteNode(root.Left, key)
	case root.Val < key:
		root.Right = deleteNode(root.Right, key)
	case root.Left == nil || root.Right == nil:
		//左右子节点有一个为空，直接判断哪个为空，就返回哪个子节点
		if root.Left != nil {
			return root.Left
		}
		return root.Right
	default:
		//默认情况就是 左右子节点都不为空的情况
		//找到右子树最小的那个节点，将其赋值给当前节点，然后递归删除那个最小的右节点
		successor := root.Right
		for successor.Left != nil {
			successor = successor.Left
		}
		successor.Right = deleteNode(root.Right, successor.Val)
		successor.Left = root.Left
		return successor
	}
	return root
}
