package linkedList

type ListNode struct {
	Val  int
	Next *ListNode
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	ans := &ListNode{Val: 0}
	//标记根节点
	i := ans
	pre := ans
	for l1 != nil || l2 != nil {
		if ans == nil {
			ans = &ListNode{Val: 0}
			pre.Next = ans
		}

		sum := ans.Val
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		ans.Val = sum % 10

		if sum >= 10 {
			ans.Next = &ListNode{Val: 0}
			ans.Next.Val += 1
		}
		pre = ans
		ans = ans.Next

	}

	return i
}
