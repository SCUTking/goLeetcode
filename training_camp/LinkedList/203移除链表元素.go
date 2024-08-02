package LinkedList

func removeElements(head *ListNode, val int) *ListNode {

	pre := &ListNode{Next: head}
	res := pre.Next
	for res.Next != nil {
		if res.Next.Val == val {
			res.Next = res.Next.Next
		} else {
			res = res.Next
		}
	}

	return pre.Next

	//temp := head
	//pre := &ListNode{Next: head}
	//for temp.Next != nil {
	//	if temp.Val == val {
	//		pre.Next = temp.Next
	//		temp = nil
	//	}
	//	pre = pre.Next
	//	temp = temp.Next
	//}
	//
	//return head

}
