package LinkedList

func swapPairs(head *ListNode) *ListNode {

	tempNode := &ListNode{Next: head}
	pre := tempNode
	temp := head
	for temp != nil || temp.Next != nil {
		node3 := temp.Next.Next
		node2 := temp.Next
		node1 := temp
		node1.Next = node3
		node2.Next = node1
		pre.Next = node2

		//下一个

		pre = temp
		temp = node3
	}

	return tempNode.Next
}
