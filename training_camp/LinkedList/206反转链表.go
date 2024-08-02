package LinkedList

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	tempNode := &ListNode{Next: head}
	flag := head
	for head != nil {
		temp := head.Next
		head.Next = tempNode
		tempNode = head
		head = temp
	}
	flag.Next = nil
	return tempNode

}
func reverseList1(head *ListNode) *ListNode {
	tempNode := &ListNode{}
	flag := head
	for flag != nil {
		temp := flag.Next
		flag.Next = tempNode
		tempNode = flag
		flag = temp
	}
	return tempNode
}
func reverseList2(head *ListNode) *ListNode {
	if head != nil || head.Next != nil {
		return head
	}
	newHead := reverseList(head)
	head.Next.Next = head
	head = head.Next

	return newHead
}
