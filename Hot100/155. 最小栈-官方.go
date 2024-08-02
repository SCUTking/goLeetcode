package Hot100

type MinStack struct {
	stack    []int
	minStack []int //标记每个Index对应的最小的值
}

func Constructor() MinStack {
	return MinStack{
		stack:    make([]int, 0),
		minStack: make([]int, 0),
	}
}

func (this *MinStack) Push(val int) {

	if len(this.stack) > 0 {
		minLast := this.minStack[len(this.minStack)-1]
		if minLast > val {
			this.minStack = append(this.minStack, val)
		} else {
			this.minStack = append(this.minStack, minLast)
		}
	} else {
		this.minStack = append(this.minStack, val)
	}
	this.stack = append(this.stack, val)
}

func (this *MinStack) Pop() {
	index := len(this.stack) - 1
	this.stack = this.stack[:index]
	this.minStack = this.minStack[:index]

}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.minStack[len(this.minStack)-1]
}
