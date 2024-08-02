package stack_and_queue

// 小根堆
type myHeap struct {
	arr []int //索引为0不操作
}

// 下滤
func (myHeap *myHeap) percolateDown(hole int) {
	currentSize := len(myHeap.arr) - 1
	var child int
	tmp := myHeap.arr[hole]
	for hole*2 <= currentSize {
		child = hole * 2
		if child != currentSize && myHeap.arr[child+1] < myHeap.arr[child] {
			child++
		}
		if myHeap.arr[child] < tmp {
			myHeap.arr[hole] = myHeap.arr[child]
		} else {
			break
		}
		hole = child
	}
	myHeap.arr[hole] = tmp
}

func (myHeap *myHeap) insert(new int) {

}
func (myHeap *myHeap) deleteMin() int {
	if myHeap.isEmpty() {
		panic("堆的大小不够")
	}
	min := myHeap.arr[1]
	currentSize := len(myHeap.arr) - 1
	myHeap.arr[1] = myHeap.arr[currentSize]
	myHeap.percolateDown(1)
	return min
}

func (myHeap *myHeap) isEmpty() bool {
	return len(myHeap.arr) < 1
}

func (myHeap *myHeap) getMin() int {
	return myHeap.arr[1]
}
func (myHeap *myHeap) buildmyHeap() {
	currentSize := len(myHeap.arr) - 1
	for i := currentSize / 2; i > 0; i-- {
		myHeap.percolateDown(i)
	}
}

func (myHeap *myHeap) buildBySlice(nums []int) {
	myHeap.arr = make([]int, len(nums)+1)
	for i, num := range nums {
		myHeap.arr[i+1] = num
	}
	myHeap.buildmyHeap()
}
