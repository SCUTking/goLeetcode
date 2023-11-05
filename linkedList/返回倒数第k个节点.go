package linkedList

func kthToLast(head *ListNode, k int) int {
	gap := 0
	slow, fast := head, head

	for fast.Next != nil {
		if gap < k-1 {
			fast = fast.Next
			gap++
		} else {
			slow = slow.Next
			fast = fast.Next
		}
	}
	return slow.Val
}
