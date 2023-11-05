package random1

func findChampion(grid [][]int) int {
	type node struct {
		d        int
		index    int
		winNodes []int
	}
	nodeList := []node{}
	for i := 0; i < len(grid); i++ {
		nodeTmp := node{d: 0, index: i, winNodes: []int{}}
		for i2, i3 := range grid[i] {
			if i3 == 0 && i2 != i {
				nodeTmp.d++
				nodeTmp.winNodes = append(nodeTmp.winNodes, i2)
			}
		}
		nodeList = append(nodeList, nodeTmp)
	}
	var queue []node
	for _, n := range nodeList {
		if n.d == 0 {
			queue = append(queue, n)
		}
	}
	cur := node{}
	for len(queue) > 0 {
		cur = queue[0]
		queue = queue[1:]
		for _, winNode := range cur.winNodes {
			nodeList[winNode].d--
			if nodeList[winNode].d == 0 {
				queue = append(queue, nodeList[winNode])
			}
		}
	}

	return cur.index

}
