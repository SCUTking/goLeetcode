package linkedList

// 算法要点：先获取整个链表的长度，再从前往后进行遍历
// 时间复杂度：O（n）
// 空间复杂度：O（1）
// 未能考虑到的点：
// 利用快慢指针，先让快指针移动n步，再使用同时移动快慢指针，这样可以只遍历一次。不用像这样遍历两次
func removeNthFromEndNode(head *ListNode, n int) *ListNode {
	L := 0
	temp := head
	for temp != nil {
		L++
		temp = temp.Next
	}

	count := 0
	index := L - n + 1
	pre := &ListNode{Next: head}
	res := pre

	for head != nil {
		count++
		if count == index {
			pre.Next = head.Next
			break
		}
		head = head.Next
		pre = pre.Next
	}

	return res.Next
}

// 遍历一次
func removeNthFromEndNode1(head *ListNode, n int) *ListNode {
	pre := &ListNode{Next: head}
	res := pre

	fast, slow := head, head
	for i := 0; i < n; i++ {
		fast = fast.Next
	}

	for fast != nil {
		fast = fast.Next
		slow = slow.Next
		pre = pre.Next
	}

	pre.Next = slow.Next

	return res.Next
}
