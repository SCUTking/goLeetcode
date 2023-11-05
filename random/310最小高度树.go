package random

import "math"

func FindMinHeightTrees(n int, edges [][]int) []int {
	type edge struct {
		Root    int
		SonNode []int
	}
	m := make(map[int]edge)

	for _, edge1 := range edges {
		i1 := edge1[0]
		i2 := edge1[1]

		e, ok := m[i1]
		if ok {
			node := e.SonNode
			node = append(node, i2)
			e.SonNode = node
			m[i1] = e
		} else {
			e2 := edge{Root: i1, SonNode: []int{i2}}
			m[i1] = e2
		}

		e1, ok1 := m[i2]
		if ok1 {
			node := e1.SonNode
			node = append(node, i1)
			e.SonNode = node
			m[i1] = e
		} else {
			e2 := edge{Root: i2, SonNode: []int{i1}}
			m[i2] = e2
		}

	}

	var mintree func(root int, used []int) int
	mintree = func(root int, used []int) int {
		if len(used) == n {
			return 0
		}
		sonNode := m[root].SonNode
		uints := minus(sonNode, used)

		if len(uints) == 0 {
			return 1
		}
		min := math.MaxInt
		for _, i := range uints {
			used := append(used, root)
			i2 := mintree(i, used) + 1
			if i2 < min {
				min = i2
			}
		}
		return min
	}

	m2 := make(map[int]int)
	ans := math.MaxInt32
	for i := 0; i < n; i++ {
		i2 := mintree(i, []int{})
		m2[i] = i2
		if i2 < ans {
			ans = i2
		}
	}

	res := []int{}
	for k, i := range m2 {
		if i == ans {
			res = append(res, k)
		}

	}

	return res
}

// minus 获取差集
func minus(a []int, b []int) []int {
	var inter []int
	mp := make(map[int]bool)
	for _, s := range a {
		if _, ok := mp[s]; !ok {
			mp[s] = true
		}
	}
	for _, s := range b {
		if _, ok := mp[s]; ok {
			delete(mp, s)
		}
	}
	for key := range mp {
		inter = append(inter, key)
	}
	return inter

}
