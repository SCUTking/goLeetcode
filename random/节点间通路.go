package random

func FindWhetherExistsPath(n int, graph [][]int, start int, target int) bool {
	edges := make([][]int, n)

	for _, e := range graph {
		i1 := e[0]
		i2 := e[1]
		edges[i1] = append(edges[i1], i2)
		edges[i2] = append(edges[i2], i1)
	}

	visited := make(map[int]bool)
	var dfs func(start int, target int) bool
	dfs = func(start int, target int) bool {
		_, ok := visited[start]
		if ok {
			return false
		} else {
			visited[start] = true

			for _, i := range edges[start] {
				if i == target {
					return true
				}
				b := dfs(i, target)
				if b {
					return true
				}
			}
		}
		return false
	}
	return dfs(start, target)
}
