package random1

func findChampion2(n int, edges [][]int) int {
	n = len(edges)
	type node struct {
		d        int
		index    int
		winNodes []int
	}
	rg := make([][]int, n)
	inDeg := make([]int, n)

	for _, edge := range edges {
		ints := edge[0]
		ints1 := edge[1]
		inDeg[ints1]++
		rg[ints] = append(rg[ints], ints1)
	}

	var queue []int
	for i, n := range inDeg {
		if n == 0 {
			queue = append(queue, i)
		}
	}

	//变更
	//if len(queue) == 1 {
	//	return queue[0]
	//} else {
	//	return -1
	if len(queue) == 1 {
		return queue[0]
	} else {
		return -1
	}

}
