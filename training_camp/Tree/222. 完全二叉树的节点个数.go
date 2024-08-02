package Tree

import "sort"

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	level := 0
	for node := root; node.Left != nil; node = node.Left {
		level++
	}

	return sort.Search(1<<(level+1), func(k int) bool {
		if k <= 1<<level {
			return false
		}
		bits := 1 << (level - 1)
		node := root
		for node != nil && bits > 0 {
			if bits&k == 0 {
				node = node.Left
			} else {
				node = node.Right
			}
			bits >>= 1
		}
		return node == nil
	}) - 1
}

func MyCountNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}

	level := 0
	node := root
	for node != nil {
		level++
		node = node.Left
	}

	return sort.Search(1<<(level), func(k int) bool {
		if k <= 1<<(level-1) {
			return false
		}

		bits := 1 << (level - 2)
		node1 := root
		for bits > 0 && node1 != nil {
			if k&bits == 0 {
				node1 = node1.Left
			} else {
				node1 = node1.Right
			}
			bits >>= 1
		}

		return node == nil

	}) - 1

}
