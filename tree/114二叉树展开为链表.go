package tree

import "fmt"

func Flatten(root *TreeNode) {

	var arr []int
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {

		if root == nil {
			return
		}
		arr = append(arr, root.Val)

		dfs(root.Left)
		dfs(root.Right)

	}
	dfs(root)

	if len(arr) <= 1 {
		return
	}
	arr = arr[1:]

	var createLinkedList func(root *TreeNode)
	createLinkedList = func(root *TreeNode) {
		if len(arr) <= 0 {
			return
		}
		first := arr[0]
		arr = arr[1:]
		node := TreeNode{
			Val: first,
		}
		root.Left = nil
		root.Right = &node
		createLinkedList(&node)
	}

	//var preorderTraversal func(root *TreeNode) ([]*TreeNode)
	//
	//preorderTraversal = func(root *TreeNode) []*TreeNode {
	//	list := []*TreeNode{}
	//	if root != nil {
	//		list = append(list, root)
	//		list = append(list, preorderTraversal(root.Left)...)
	//		list = append(list, preorderTraversal(root.Right)...)
	//	}
	//	return list
	//}
	//
	//list := preorderTraversal(root)
	//for i := 1; i < len(list); i++ {
	//	prev, curr := list[i-1], list[i]
	//	prev.Left, prev.Right = nil, curr
	//}

	createLinkedList(root)

}

func Flatten1(root *TreeNode) {
	list := preorderTraversal(root)
	for i := 1; i < len(list); i++ {
		prev, curr := list[i-1], list[i]
		prev.Left, prev.Right = nil, curr
	}
	fmt.Println(root)
}

func preorderTraversal(root *TreeNode) []*TreeNode {
	list := []*TreeNode{}
	if root != nil {
		list = append(list, root)
		list = append(list, preorderTraversal(root.Left)...)
		list = append(list, preorderTraversal(root.Right)...)
	}
	return list
}
