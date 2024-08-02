package LinkedList

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	tempNode := &ListNode{Next: head}
	temp := tempNode
	sum := 0
	for temp.Next != nil {
		sum++
		temp = temp.Next
	}

	count := 0
	temp = tempNode
	for temp.Next != nil {

		if count == sum-n {
			temp.Next = temp.Next.Next
			break
		}
		count++
	}

	return tempNode.Next

}

func removeNthFromEnd1(head *ListNode, n int) *ListNode {
	nodes := []*ListNode{}
	dummy := &ListNode{0, head}
	for node := dummy; node != nil; node = node.Next {
		nodes = append(nodes, node)
	}
	prev := nodes[len(nodes)-1-n]
	prev.Next = prev.Next.Next
	return dummy.Next

}
