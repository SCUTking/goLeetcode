package unionFindSet

func FindRedundantConnection(edges [][]int) []int {

	ufSet := make(map[int]int, len(edges))

	var find func(index int) int
	var union func(index int, index1 int)

	find = func(index int) int {
		fa, exist := ufSet[index]
		if !exist {
			ufSet[index] = index
			return index
		}
		if fa == index {
			return index
		}
		//如果不是本身的话，进行路径压缩，将树的高度减少
		fa = find(fa)
		ufSet[index] = fa
		return fa
	}

	//合并就是根指针指向合并罢了
	union = func(index int, index1 int) {
		fa1 := ufSet[index]
		fa2 := ufSet[index1]
		if fa2 != fa1 {
			//有个“小挂大”的优化，这里就没有优化
			ufSet[fa1] = fa2
		}
	}

	for _, edge := range edges {
		i := edge[0]
		i1 := edge[1]

		//说明不是同一个集和
		if find(i) != find(i1) {
			union(i, i1)
		} else {
			return edge
		}

	}

	return nil

}
