package linkedList

// 算法要点：利用map能在O（1）内判断节点是否存在，如果有相交的节点，就会出现map命中
// 时间复杂度：O（n）
// 空间复杂度：O（n）
// 未能考虑到的点：
//	如果不使用map，可以计算两链表的长度，求出差值，将两链表的后部分进行对齐

func LinkedCross(headA, headB *ListNode) *ListNode {
	m := make(map[*ListNode]bool, 0)
	for headA != nil {
		m[headA] = true
		headA = headA.Next
	}
	for headB != nil {
		if m[headB] {
			return headB
		}
		headB = headB.Next
	}
	return nil
}
