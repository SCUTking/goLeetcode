package linkedList

// 算法要点：循环条件要加上head.Next也不为nil，因为要凑够两个的才能进行交换
// 时间复杂度：O（n）
// 空间复杂度：O（1）
// 未能考虑到的点：

func swapListNode(head *ListNode) *ListNode {
	pre := &ListNode{Next: head}
	res := pre
	for head != nil && head.Next != nil {
		lastNode := head.Next.Next
		pre.Next = head.Next
		head.Next.Next = head
		head.Next = lastNode
		pre = head
		head = lastNode
	}
	return res.Next

}
