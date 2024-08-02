package linkedList

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil { //判断是否只有一个节点：根据后面是否是还有节点判断
		return head
	}

	//利用快慢指针找到中间节点
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	mid := slow.Next //之所以用Next，是因为用slow话，不好进行截断
	slow.Next = nil  //截断，将链表分割为两部分处理

	LeftList := sortList(head)
	RightList := sortList(mid)

	return MergeLinkedList(LeftList, RightList)
}

// MergeLinkedList 合并两有序链表
func MergeLinkedList(l, r *ListNode) *ListNode {
	if l == nil {
		return r
	}
	if r == nil {
		return l
	}
	temp := &ListNode{Next: nil}
	res := temp //标记要返回的节点
	lp, rp := l, r
	for lp != nil && rp != nil { //不用Next，因为本身就要对比
		if lp.Val < rp.Val {
			temp.Next = lp
			lp = lp.Next
		} else {
			temp.Next = rp
			rp = rp.Next
		}
		temp = temp.Next
	}

	for lp != nil {
		temp.Next = lp
		temp = temp.Next
		lp = lp.Next
	}

	for rp != nil {
		temp.Next = rp
		temp = temp.Next
		rp = rp.Next
	}

	return res.Next
}
