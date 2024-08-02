package Hot100

import "math"

//首先，定义一个数组 D，D[v] 表示从源点 s 到顶点 v 的边的权值，如果没有边则将 D[v] 置为无穷大。
//把图的顶点集合划分为两个集合 S 和 V-S。第一个集合 S 表示距源点最短距离已经确定的顶点集，即一个顶点如果属于集合 S 则说明从源点 s 到该顶点的最短路径已知。其余的顶点放在另一个集合 V-S 中。
//每次从尚未确定最短路径长度的集合 V-S 中取出一个最短特殊路径长度最小的顶点 u（这个路径是从起点经过任意集合S节点到达的最小节点），将 u 加入集合 S，同时修改数组 D 中由 s 可达的最短路径长度。若加入集合 S 的 u 作为中间顶点时，vi 的最短路特殊路径长度变短，则修改 vi 的距离值（即当D[u] + W[u, vi] < D[vi]时，令D[vi] = D[u] + W[u, vi]）。
//重复第 3 步的操作，一旦 S 包含了所有 V 中的顶点，D 中各顶点的距离值就记录了从源点 s 到该顶点的最短路径长度。

func dijkstra(graph [][]int, startNode int) []int {
	//获取节点数
	n := len(graph)

	visited := make([]bool, n) //标记是否访问过，利用布尔数组好一些，不要用集合
	distance := make([]int, n) //记录到各个节点距离的

	//初始化
	for i := 0; i < n; i++ {
		visited[i] = false          //初始没访问过
		distance[i] = math.MaxInt32 //初始化是无穷大，不可达
	}
	//起始节点，标记为0，就可以在循环中处理，不用再对起始节点进行处理
	distance[startNode] = 0
	//通过是否还能找到下一个最小遍历节点来判断循环是否结束
	for {
		minPath := math.MaxInt32
		curNode := -1

		//利用遍历o(n)找寻（距离已遍历节点集合）最近的节点
		for j := 0; j < n; j++ {
			if !visited[j] && minPath > distance[j] {
				curNode = j
				minPath = distance[j]
			}
		}

		//如果没有节点便利了，退出循环
		if curNode == -1 {
			break
		}

		//标记已经访问
		visited[curNode] = true

		for j := 0; j < n; j++ {
			//排除不可达的 注意-1可能会变成别的
			if graph[curNode][j] != -1 {
				//新路径：到达curNode最小路径的+该节点到curNode
				newPath := minPath + graph[curNode][j]
				if !visited[j] && newPath < distance[j] {
					distance[j] = newPath
				}
			}
		}
	}

	return distance
}
