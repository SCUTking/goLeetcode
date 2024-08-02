package main

import (
	"fmt"
)

type node struct {
	parent *node
}

func main() {

	//利用并查集
	n := 0
	m := 0
	q := 0
	fmt.Scan(&n, &m, &q)
	bind := make([]*node, n)
	for i := 0; i < n; i++ {
		bind[i] = &node{parent: nil}
	}
	for m > 0 {
		m--
		var a, b int
		fmt.Scan(&a, &b)
		bind[a-1].parent = bind[b-1]
	}

	for q > 0 {
		q--
		var a, b, c int
		fmt.Scan(&a, &b, &c)
		if a == 1 {
			if bind[b-1].parent == bind[c-1] {
				bind[b-1] = nil
			}
			if bind[c-1].parent == bind[b-1] {
				bind[c-1] = nil
			}
		} else {

			fmt.Println("Yes")
			fmt.Println("No")
		}
	}
}
