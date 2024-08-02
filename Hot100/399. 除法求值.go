package Hot100

type node struct {
	val       string
	childs    []*node
	childVals []float64
}

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	m := make(map[string]*node)
	for i := 0; i < len(equations); i++ {
		p := equations[i][0]
		c := equations[i][1]
		v := values[i]
		if _, e := m[p]; !e {
			m[p] = NewNode(p)
		}
		if _, e := m[c]; !e {
			m[c] = NewNode(c)
		}

		m[p].childs = append(m[p].childs, m[c])
		m[p].childVals = append(m[p].childVals, v)
		m[c].childs = append(m[c].childs, m[p])
		m[c].childVals = append(m[c].childVals, 1/v)
	}

	perRes := -1.0
	visited := make(map[string]bool, 0)
	var dfs func(root *node, target string, pre float64)
	dfs = func(root *node, target string, pre float64) {
		if visited[root.val] {
			return
		}
		if root.val == target {
			perRes = pre
			return
		}
		visited[root.val] = true
		for i := 0; i < len(root.childs); i++ {
			dfs(root.childs[i], target, pre*root.childVals[i])
		}
	}

	res := make([]float64, len(queries))
	for i := 0; i < len(queries); i++ {
		p := queries[i][0]
		c := queries[i][1]
		if m[p] == nil || m[c] == nil {
			res[i] = -1.0
			continue
		}
		visited = make(map[string]bool)
		dfs(m[p], m[c].val, 1.0)

		res[i] = perRes
		perRes = -1.0
	}

	return res

}

func NewNode(s string) *node {
	return &node{
		val:       s,
		childs:    make([]*node, 0),
		childVals: make([]float64, 0),
	}
}
