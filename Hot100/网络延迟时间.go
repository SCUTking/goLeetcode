package Hot100

import "math"

func networkDelayTime(times [][]int, n int, k int) int {
	graph := make([][]int, n)
	for i := 0; i < n; i++ {
		graph[i] = make([]int, n)
		for j := 0; j < n; j++ {
			graph[i][j] = -1
		}
	}
	for i := 0; i < len(times); i++ {
		x := times[i][0]
		y := times[i][1]
		graph[x][y] = times[i][2]
	}

	ints := dijkstra(graph, k)
	res := -1
	for i := 0; i < len(ints); i++ {
		if ints[i] == math.MaxInt32 {
			return -1
		}
		if res < ints[i] {
			res = ints[i]
		}
	}

	return res

}
