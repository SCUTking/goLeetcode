package Hot100

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	m := make(map[*ListNode]struct{}, 0)
	for headA != nil {
		m[headA] = struct{}{}
		headA = headA.Next
	}

	for headB != nil {
		_, exist1 := m[headB]
		if exist1 {
			return headB
		}
		headB = headB.Next
	}

	return nil
}
