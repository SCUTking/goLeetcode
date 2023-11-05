package linkedList

func RotateRight(head *ListNode, k int) *ListNode {
	size := GetLinkedListSize(head)
	if size == 0 {
		return head
	}
	k = k % size
	frist := head
	for i := 0; i < k; i++ {
		frist = moveRightOne(frist)
	}

	return frist
}

func moveRightOne(head *ListNode) *ListNode {
	pre := &ListNode{Next: head}
	for pre.Next.Next != nil {
		pre = pre.Next
	}
	last := pre.Next
	last.Next = head
	pre.Next = nil
	return last
}

func GetLinkedListSize(head *ListNode) int {
	if head == nil {
		return 0
	}
	n := 1
	for head.Next != nil {
		n++
		head = head.Next
	}
	return n
}
