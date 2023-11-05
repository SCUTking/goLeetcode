package random1

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func ReorderList(head *ListNode) {
	if head.Next == nil {
		return
	}
	var getLast func(head *ListNode) *ListNode
	getLast = func(head *ListNode) *ListNode {
		temp := head
		for temp.Next.Next != nil {
			temp = temp.Next
		}
		return temp
	}

	left := head

	for left != nil && left.Next != nil {
		secondlast := getLast(left)
		if secondlast != left {
			secondlast.Next.Next = left.Next
		}
		left.Next = secondlast.Next
		if secondlast != left {
			secondlast.Next = nil
		}

		left = left.Next.Next
	}
}
