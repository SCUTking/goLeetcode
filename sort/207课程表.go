package sort

func CanFinish(numCourses int, prerequisites [][]int) bool {

	edges := [][]int{}
	queue := []int{}
	d := []int{}
	for i := 0; i < numCourses; i++ {
		d = append(d, 0)
		edges = append(edges, []int{})
	}

	for _, prerequisite := range prerequisites {
		pre := prerequisite[0]
		post := prerequisite[1]
		d[post]++

		ints := edges[pre]
		ints = append(ints, post)
		edges[pre] = ints

	}

	for i, v := range d {
		if v == 0 {
			queue = append(queue, i)
		}
	}

	for len(queue) > 0 {
		for _, i := range queue {
			queue = queue[1:]
			for _, i3 := range edges[i] {
				d[i3]--
				if d[i3] == 0 {
					queue = append(queue, i3)
				}
			}
			numCourses--
		}
	}
	if numCourses > 0 {
		return false
	}
	return true

}
