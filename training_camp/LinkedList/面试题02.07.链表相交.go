package LinkedList

func getIntersectionNode(headA, headB *ListNode) *ListNode {

	a := make([]*ListNode, 0)

	for headA != nil {
		a = append(a, headA)
		headA = headA.Next
	}
	for headB != nil {
		for _, node := range a {
			if node == headB {
				return headB
			}
		}
		headB = headB.Next
	}
	return nil
}
func getIntersectionNode1(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	pa, pb := headA, headB
	for pa != pb {
		if pa == nil {
			pa = headB
		} else {
			pa = pa.Next
		}
		if pb == nil {
			pb = headA
		} else {
			pb = pb.Next
		}
	}

	return pa
}
