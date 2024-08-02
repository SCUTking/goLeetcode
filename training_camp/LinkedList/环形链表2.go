package LinkedList

func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head
	temp := head
	for fast != nil {
		slow = slow.Next
		if fast.Next == nil {
			return nil
		}
		fast = fast.Next.Next

		if fast == slow {
			for slow != temp {
				slow = slow.Next
				temp = temp.Next
			}
			return slow
		}
	}
	return nil
}
