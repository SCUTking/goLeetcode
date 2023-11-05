package Tric

import "strconv"

func LexicalOrder(n int) []int {

	root := &TreeNode{}
	for i := 1; i <= n; i++ {
		root.AddString(i)
	}

	var res []int
	//层序遍历
	var dfs func(root *TreeNode, qianzhui string, index int)
	dfs = func(root *TreeNode, qianzhui string, index int) {
		if root == nil {
			return
		}
		if root != nil {
			num, _ := strconv.Atoi(qianzhui + strconv.Itoa(index))
			res = append(res, num)
		}
		for i, _ := range root.next {
			dfs(root.next[i], qianzhui+strconv.Itoa(index), i)
		}
	}

	for i, node := range root.next {
		if node != nil {
			dfs(root.next[i], "", i)
		}
	}

	return res
}

type TreeNode struct {
	next [10]*TreeNode
}

func (this *TreeNode) AddString(num int) {
	itoa := strconv.Itoa(num)
	nowNode := this
	for _, i := range itoa {
		iNum := i - '0'
		if nowNode.next[iNum] == nil {
			nowNode.next[iNum] = &TreeNode{}
		}
		nowNode = nowNode.next[iNum]
	}
}
