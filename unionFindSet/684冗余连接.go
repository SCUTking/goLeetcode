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
func FindRedundantConnection1(edges [][]int) []int {
	parent := make([]int, len(edges)+1)
	//初始化时，父节点指向自己
	for i, _ := range parent {
		parent[i] = i
	}

	//find 查找某个节点的根节点
	//可用于判断是否在同一个集合中 判断条件：find(a)==find(b)
	//
	var find func(node int) int
	find = func(node int) int {
		if parent[node] == node {
			return node
		}

		return find(parent[node])
	}

	//union 判断是否能够合并
	//如果能够合并  返回成功
	//如果不能合并  返回false
	var union func(a int, b int) bool
	union = func(a int, b int) bool {
		x := find(a)
		y := find(b)
		if x == y {
			return false
		}
		//直接父的根节点相互指向即可
		parent[x] = y
		return true
	}

	for i := 0; i < len(edges); i++ {
		if !union(edges[i][0], edges[i][1]) {
			return edges[i]
		}
	}

	return nil
}
