package stack_and_queue

import "container/heap"

func topKFrequent(nums []int, k int) []int {
	m := make(map[int]int, 0)
	for _, num := range nums {
		m[num]++
	}

	stack := make([]int, 0)
	for _, num := range m {
		temp := make([]int, 0)
		for len(stack) > 0 && stack[len(stack)-1] > num {
			temp = append(temp, stack[len(stack)-1])
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, num)
		for i := len(temp) - 1; i >= 0; i-- {
			stack = append(stack, temp[i])
		}
	}
	res := make([]int, 0)
	n := len(stack) - 1
	for i := n; i >= n-k+1; i-- {
		res = append(res, stack[i])
	}
	return res
}
func topKFrequent1(nums []int, k int) []int {
	occurrences := map[int]int{}
	for _, num := range nums {
		occurrences[num]++
	}
	h := &IHeap{}
	heap.Init(h)
	for key, value := range occurrences {
		heap.Push(h, [2]int{key, value})
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	ret := make([]int, k)
	for i := 0; i < k; i++ {
		ret[k-i-1] = heap.Pop(h).([2]int)[0]
	}
	return ret
}

type IHeap [][2]int

func (h IHeap) Len() int           { return len(h) }
func (h IHeap) Less(i, j int) bool { return h[i][1] < h[j][1] }
func (h IHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

// 放到最后，然后内部会帮我们上滤
func (h *IHeap) Push(x interface{}) {
	*h = append(*h, x.([2]int))
}

// Pop 最后一个元素，内部方法会将第一个元素先与最后一个元素交换的，所以直接返回最后一个元素即可
func (h *IHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
