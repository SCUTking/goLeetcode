package linkedList

// 算法要点：利用前置节点，调整前后指向
// 时间复杂度：O（n）
// 空间复杂度：O（1）
// 未能考虑到的点：
// 前置节点不能指向Head，可能会出现环
// 同时需要考虑的head为nil的情况
// 前置节点不应该设置下一个节点
func reverseLinkedList(head *ListNode) *ListNode {
	var pre *ListNode
	for head != nil {
		nextNode := head.Next
		head.Next = pre
		pre = head
		head = nextNode
	}
	return pre
}
