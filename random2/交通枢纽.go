package random2

type node struct {
	num  int
	rudu int
	next []int
}

func transportationHub(path [][]int) int {
	nodes := make(map[int]struct{})
	outDeg := make(map[int]int)
	inDeg := make(map[int]int)

	for _, ints := range path {
		i := ints[0]
		i2 := ints[1]
		nodes[i] = struct{}{}
		nodes[i2] = struct{}{}
		outDeg[i]++
		inDeg[i2]--
	}

	for i, _ := range nodes {
		if outDeg[i] == 0 && inDeg[i] == len(path)-1 {
			return i
		}
	}

	return -1

}

func my_transportationHub(path [][]int) int {
	nodes := make(map[int]*node, 0)
	for _, ints := range path {
		sourse := ints[0]
		target := ints[1]

		_, ok := nodes[sourse]
		if !ok {
			nodes[sourse] = &node{
				num:  sourse,
				rudu: 0,
				next: make([]int, 0),
			}
		}

		_, ok1 := nodes[target]
		if !ok1 {
			nodes[target] = &node{
				num:  target,
				rudu: 0,
				next: make([]int, 0),
			}
		}
		nodes[target].rudu++
		nodes[sourse].next = append(nodes[sourse].next, target)
	}
	oldNum := len(nodes)
	queue := make([]*node, 0)
	for _, node := range nodes {
		if len(node.next) == 0 && node.rudu == oldNum-1 {
			return node.num
		}
		if node.rudu == 0 {
			queue = append(queue, node)
		}
	}
	return -1

}
