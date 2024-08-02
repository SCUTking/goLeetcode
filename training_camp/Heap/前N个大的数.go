package Heap

type myHeap []int

func (h myHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h myHeap) Len() int {
	return len(h)
}
func (h myHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *myHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *myHeap) Pop() interface{} {

	old := *h
	n := len(old)
	res := old[n-1]
	*h = old[:n-1]
	return res
}
