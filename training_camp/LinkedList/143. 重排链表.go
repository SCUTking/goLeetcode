package LinkedList

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList(head *ListNode) {
	h := []*ListNode{}
	temp := head
	for temp != nil {
		h = append(h, temp)
		temp = temp.Next
	}
	n := len(h)
	temp = head
	for i := n; i > n/2; i-- {
		node := temp.Next
		h[i-1].Next = node
		temp.Next = h[i-1]
		temp = node
	}

}
