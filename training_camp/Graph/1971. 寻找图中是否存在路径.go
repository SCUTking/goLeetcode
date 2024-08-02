package Graph

func validPath(n int, edges [][]int, source int, destination int) bool {
	parent := make([]int, n)
	//初始化
	for i := 0; i < n; i++ {
		parent[i] = i
	}

	var find func(child int) int
	find = func(child int) int {
		if parent[child] == child {
			return child
		}
		parent[child] = find(parent[child]) //路径压缩
		return parent[child]
	}

	var join func(i, j int)
	join = func(i, j int) {
		u := find(i)
		v := find(j)
		if u == v {
			return
		}
		parent[u] = v
	}

	var isSame func(i, j int) bool
	isSame = func(i, j int) bool {
		u := find(i)
		v := find(j)
		if u == v {
			return true
		}
		return false
	}

	for i := 0; i < len(edges); i++ {
		x := edges[i][0]
		y := edges[i][1]
		join(x, y)
	}

	return isSame(source, destination)
}
