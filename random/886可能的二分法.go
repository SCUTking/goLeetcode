package random

func PossibleBipartition(n int, dislikes [][]int) bool {
	findSet := []int{}
	for i := 0; i < n; i++ {
		findSet = append(findSet, i)
	}

	var find func(n int) int
	find = func(n int) int {
		if findSet[n] != n {
			//说明父节点不是自己  所以不是根节点
			return find(findSet[n])
		}
		return n
	}

	var union func(a, b int)
	union = func(a, b int) {
		if find(a) == find(b) {
			return
		}
		if find(a) < find(b) {
			findSet[a] = b
		} else {
			findSet[b] = a
		}

	}

	edges := make([][]int, n)
	for _, dislike := range dislikes {
		i1 := dislike[0] - 1
		i2 := dislike[1] - 1
		edges[i1] = append(edges[i1], i2)
		edges[i2] = append(edges[i2], i1)
	}

	for index, edge := range edges {

		for _, i := range edge {
			union(edge[0], i)
			if find(index) == find(i) {
				return false
			}
		}

	}

	return true

}

//func PossibleBipartition(n int, dislikes [][]int) bool {
//	edges := make([][]int, n)
//	for _, dislike := range dislikes {
//		i1 := dislike[0] - 1
//		i2 := dislike[1] - 1
//		edges[i1] = append(edges[i1], i2)
//		edges[i2] = append(edges[i2], i1)
//	}
//	mapNInWhich := make(map[int]int)
//	if len(dislikes) == 0 {
//		return true
//	}
//	for i, edge := range edges {
//		mapNInWhich[i] = 0
//		if len(edge) > 0 {
//			mapNInWhich[i] = 1
//			for _, i2 := range edge {
//				mapNInWhich[i2] = 2
//			}
//			break
//		}
//	}
//
//	for i, edge := range edges {
//		if mapNInWhich[i] == 0 {
//			for _, i2 := range edge {
//				if mapNInWhich[i2] != 0 {
//					if mapNInWhich[i2] == 1 {
//						mapNInWhich[i] = 2
//					} else {
//						mapNInWhich[i] = 1
//					}
//					break
//				}
//			}
//
//		}
//		if mapNInWhich[i] == 1 {
//			for _, i2 := range edge {
//				if mapNInWhich[i2] == 1 {
//					return false
//				}
//				mapNInWhich[i2] = 2
//			}
//		} else {
//			for _, i2 := range edge {
//				if mapNInWhich[i2] == 2 {
//					return false
//				}
//				mapNInWhich[i2] = 1
//			}
//		}
//	}
//
//	return false
//}
