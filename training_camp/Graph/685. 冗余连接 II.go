package Graph

func findRedundantDirectedConnection(edges [][]int) []int {
	n := len(edges)
	in := make([]int, n+1)

	for i := 0; i < n; i++ {
		in[edges[i][1]]++
	}

	parent := make([]int, n)

	var init func()
	init = func() {
		for i := 0; i < n; i++ {
			parent[i] = i
		}
	}

	var find func(i int) int
	find = func(i int) int {
		if parent[i] == i {
			return i
		}
		parent[i] = find(parent[i])
		return parent[i]
	}

	var join func(i, j int)
	join = func(i, j int) {
		u := find(i)
		v := find(j)
		if u != v {
			return
		}
		parent[u] = v
	}

	var isSame func(i, j int) bool
	isSame = func(i, j int) bool {
		u := find(i)
		v := find(j)
		if u != v {
			return false
		}
		return true
	}

	//第三种情况，找出环中不要的那个边
	var getRemoveEdge func() []int
	getRemoveEdge = func() []int {
		init()
		for i := 0; i < n; i++ {
			x := edges[i][0]
			y := edges[i][1]
			if isSame(x, y) {
				return edges[i]
			}
		}
		return []int{}
	}

	//入度为2时，判断那两条边是不是不能构成树
	var isTreeAfterRemoveEdge func(index int) bool
	isTreeAfterRemoveEdge = func(index int) bool {
		init()
		for i := 0; i < n; i++ {
			if i == index {
				continue
			} else if isSame(edges[i][0], edges[i][1]) {
				return false
			}
			join(edges[i][0], edges[i][1])
		}
		return true
	}

	c := make([]int, 0)
	for i := 0; i < n; i++ {
		if in[i] == 2 {
			c = append(c, i-1)
		}
	}
	res := []int{}
	for i := 0; i < len(c); i++ {
		if isTreeAfterRemoveEdge(c[i]) {
			res = edges[c[i]]
		}
	}

	if res != nil {
		return res
	}

	return getRemoveEdge()

}
