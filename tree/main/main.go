package main

import "goLeetcode/tree"

func main() {
	node4 := &tree.TreeNode{Val: 15}
	node5 := &tree.TreeNode{Val: 7}
	node3 := &tree.TreeNode{Val: 20, Left: node4, Right: node5}

	node2 := &tree.TreeNode{Val: 9}

	node1 := &tree.TreeNode{Val: 3, Left: node2, Right: node3}

	tree.IsSymmetric(node1)

}
