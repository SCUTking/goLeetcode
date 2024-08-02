package Graph

func allPathsSourceTarget(graph [][]int) [][]int {

	n := len(graph)
	res := make([][]int, 0)
	path := make([]int, 1)
	var dfs func()
	dfs = func() {
		last := path[len(path)-1]
		if last == n-1 {
			temp := make([]int, len(path))
			copy(temp, path)

			res = append(res, temp)
		}
		//有向图：遍历该节点的所有可达节点，因为有向图因此不会出现死循环
		for i := 0; i < len(graph[last]); i++ {
			path = append(path, graph[last][i])
			dfs()
			path = path[:len(path)-1]
		}
	}

	dfs()

	return res
}
