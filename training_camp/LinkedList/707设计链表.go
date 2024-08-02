package LinkedList

type MyLinkedList struct {
	head *ListNode
	len  int
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func Constructor() MyLinkedList {
	return MyLinkedList{head: nil, len: 0}
}

func (this *MyLinkedList) Get(index int) int {
	count := 0
	temp := &ListNode{Next: this.head}
	for temp.Next != nil {
		if count == index {
			return temp.Next.Val
		}
		temp = temp.Next
		count++
	}
	return -1
}

func (this *MyLinkedList) AddAtHead(val int) {
	if this.len == 0 {
		this.head = &ListNode{val, nil}
	} else {
		newNode := &ListNode{Next: this.head, Val: val}
		this.head = newNode
	}
	this.len++
}

func (this *MyLinkedList) AddAtTail(val int) {
	newNode := &ListNode{val, nil}
	if this.len == 0 {
		this.head = newNode
	} else {
		temp := this.head
		for temp.Next != nil {
			temp = temp.Next
		}
		temp.Next = newNode
	}
	this.len++
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index == this.len {
		this.AddAtTail(val)
	} else if index > this.len {
		return
	} else {
		count := 0
		temp := &ListNode{Next: this.head}
		flag := temp
		for temp.Next != nil {
			if count == index {
				newNode := &ListNode{val, temp.Next}
				temp.Next = newNode
				break
			}
			temp = temp.Next
			count++
		}
		this.head = flag.Next
		this.len++
	}
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index >= this.len {
		return
	}
	count := 0
	temp := &ListNode{Next: this.head}
	flag := temp
	for temp.Next != nil {
		if count == index {
			temp.Next = temp.Next.Next
			break
		}
		temp = temp.Next
		count++
	}
	this.head = flag.Next
	this.len--
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
