package Heap

import (
	"container/heap"
	"fmt"
	"testing"
)

func Test1(t *testing.T) {
	h := &myHeap{1, 3, 2}

	heap.Init(h)
	heap.Push(h, 8)
	count := h.Len()
	for i := 0; i < count; i++ {
		pop := heap.Pop(h)
		fmt.Println(pop)
	}

}
