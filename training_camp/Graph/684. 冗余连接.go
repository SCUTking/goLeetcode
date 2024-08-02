package Graph

func findRedundantConnection(edges [][]int) []int {
	n := len(edges)
	parent := make([]int, n+1)
	//初始化
	for i := 1; i <= n; i++ {
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

	res := []int{0, 0}
	for i := 0; i < n; i++ {
		pre := edges[i][0]
		aft := edges[i][1]
		if isSame(pre, aft) {
			res = edges[i]
		}
		join(pre, aft)

	}

	return res
}
