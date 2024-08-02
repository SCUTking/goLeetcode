package linkedList

type MyLinkedList struct {
	head *ListNode
	L    int
}

func Constructor() MyLinkedList {
	return MyLinkedList{L: 0}
}

// 算法要点：利用前置节点，调整前后指向
// 时间复杂度：O（n）
// 空间复杂度：O（1）
// 未能考虑到的点： 添加头结点与尾节点  都可以复用添加指定index节点的函数
func (this *MyLinkedList) Get(index int) int {
	head := this.head
	if this.L <= index {
		return -1
	}
	count := 0
	for head != nil {
		if count == index {
			return head.Val
		}
		count++
		head = head.Next
	}
	return -1
}

func (this *MyLinkedList) AddAtHead(val int) {
	this.AddAtIndex(0, val)
}

func (this *MyLinkedList) AddAtTail(val int) {
	this.AddAtIndex(this.L, val)
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index > this.L {
		return
	}
	this.L++
	newNode := &ListNode{Val: val}
	head := this.head
	pre := &ListNode{Next: head}
	res := pre
	for i := 0; i < index; i++ {
		head = head.Next
		pre = pre.Next
	}
	pre.Next = newNode
	newNode.Next = head
	this.head = res.Next
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index >= this.L {
		return
	}
	this.L--
	head := this.head
	pre := &ListNode{Next: head}
	res := pre
	count := 0
	for head != nil {
		if count == index {
			//删除后，pre不用动了
			pre.Next = head.Next
			this.head = res.Next
			break
		}
		count++
		head = head.Next
		pre = pre.Next

	}
}
