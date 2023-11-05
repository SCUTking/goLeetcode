package linkedList

func SwapPairs(head *ListNode) *ListNode {
	ans := head
	cur := head
	flag := true
	pre := head
	for cur != nil {
		if cur.Next != nil {
			temp := cur.Next

			cur.Next = temp.Next
			temp.Next = cur
			if flag {
				flag = false
				ans = temp
			} else {
				pre.Next = temp
			}
			pre = cur
			cur = cur.Next

		} else {
			cur = cur.Next
		}
	}
	return ans

}
