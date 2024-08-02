package linkedList

// 算法要点：
// 时间复杂度：O（n）
// 空间复杂度：O（n）
// 未能考虑到的点：
func loopLinkedList(head *ListNode) *ListNode {

	fast, slow := head, head
	for fast.Next != nil && slow != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			for slow != head {
				slow = slow.Next
				head = head.Next
			}
			return slow
		}
	}

	return nil
}

// 算法要点：
// 时间复杂度：O（n）
// 空间复杂度：O（n）
// 利用hash表，环形链表，肯定会遇到相同的元素
func loopLinkedList1(head *ListNode) *ListNode {

	m := make(map[*ListNode]bool, 0)
	for head != nil {

		if m[head] {
			return head
		}
		m[head] = true
	}

	return nil
}
