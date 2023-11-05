package hash

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	visited := make(map[*Node]*Node)
	var cg func(node *Node) *Node
	cg = func(node *Node) *Node {
		if node == nil {
			return node
		}
		if visited[node] == nil {
			clone := &Node{Val: node.Val, Neighbors: []*Node{}}
			visited[node] = clone

			for _, n := range node.Neighbors {
				clone.Neighbors = append(clone.Neighbors, cg(n))
			}

			return clone

		} else {
			return visited[node]
		}

	}

	return cg(node)
}

func (fu *Node) createNeighbors(res *Node) {
	if len(fu.Neighbors) == 0 {
		res.Val = fu.Val
	} else {
		var ns []*Node
		for _, neighbor := range fu.Neighbors {
			var temp Node
			temp.Val = neighbor.Val
			ns = append(ns, &temp)
			neighbor.createNeighbors(&temp)
		}
		res.Neighbors = ns
	}
}
