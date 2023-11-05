package main

import "goLeetcode/sort"

func main() {
	//edges := [][]int{}
	////1,4],[2,4],[3,1],[3,2
	//edges = append(edges, []int{1, 4})
	////edges = append(edges, []int{0, 1})
	//edges = append(edges, []int{2, 4})
	//edges = append(edges, []int{3, 2})
	//edges = append(edges, []int{3, 1})
	//sort.CanFinish(5, edges)

	arr := []int{5, 6, 4, 2, 4, 98}

	//sort.BubbleSort(arr)
	//sort.SelectSort(arr)

	//sort.InsertSort(arr)
	//sort.ShellSort(arr)

	sort.HeapSort(arr)
}
