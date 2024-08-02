package linkedList

type ListNode struct {
	Val  int
	Next *ListNode
}

// 算法要点：设置一个前置节点
// 时间复杂度：O（n）
// 空间复杂度：O（1）
// 未能考虑到的点：
// 删除节点后，不用再移动前置节点
func removeElement(head *ListNode, target int) *ListNode {
	pre := &ListNode{Next: head}
	res := pre
	for head != nil {
		if head.Val == target {
			//删除后，pre不用动了
			pre.Next = head.Next
		} else {
			pre = pre.Next
		}
		head = head.Next
	}
	return res.Next

}
