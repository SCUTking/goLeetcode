package myself

type MinStack struct {
	stack     []int
	sortStack []int //单调索引栈，用于删除后找到上一个最小的
}

func Constructor() MinStack {
	return MinStack{
		stack:     make([]int, 0),
		sortStack: make([]int, 0),
	}
}

func (this *MinStack) inSortStack(index int) {
	temp := make([]int, 0)

	for len(this.sortStack) > 0 && this.stack[this.sortStack[len(this.sortStack)-1]] < this.stack[index] {
		temp = append(temp, this.sortStack[len(this.sortStack)-1])
		this.sortStack = this.sortStack[:len(this.sortStack)-1]
	}

	this.sortStack = append(this.sortStack, index)
	for i := len(temp) - 1; i >= 0; i-- {
		this.sortStack = append(this.sortStack, temp[i])
	}
}

func (this *MinStack) Push(val int) {
	this.stack = append(this.stack, val)
	this.inSortStack(len(this.stack) - 1)
}

func (this *MinStack) Pop() {
	index := len(this.stack) - 1
	this.stack = this.stack[:index]
	for i := 0; i < len(this.sortStack); i++ {
		if this.sortStack[i] == index {
			this.sortStack = append(this.sortStack[:i], this.sortStack[i+1:]...)
		}
	}
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.stack[this.sortStack[len(this.sortStack)-1]]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
