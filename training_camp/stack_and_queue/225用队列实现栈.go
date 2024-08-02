package stack_and_queue

type MyStack struct {
	inQueue, outQueue []int
}

func Constructor() MyStack {
	return MyStack{}
}

func (this *MyStack) Push(x int) {
	this.inQueue = append(this.inQueue, x)
}

func (this *MyStack) in2out() {
	temp := make([]int, 0)
	for len(this.inQueue) > 0 {
		temp = append(temp, this.inQueue[0])
		this.inQueue = this.inQueue[1:]
	}
	oldOut := make([]int, len(this.outQueue))
	copy(oldOut, temp)
	this.outQueue = this.outQueue[:0]
	for i := len(temp) - 1; i >= 0; i-- {
		this.outQueue = append(this.outQueue, temp[i])
	}
	this.outQueue = append(this.outQueue, oldOut...)
}

func (this *MyStack) Pop() int {
	if len(this.outQueue) == 0 || len(this.inQueue) != 0 {
		this.in2out()
	}
	res := this.outQueue[0]
	this.outQueue = this.outQueue[1:]
	return res
}

func (this *MyStack) Top() int {
	if len(this.outQueue) == 0 {
		this.in2out()
	}
	res := this.outQueue[0]
	return res
}

func (this *MyStack) Empty() bool {
	if len(this.inQueue) == 0 && len(this.outQueue) == 0 {
		return true
	}
	return false
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
