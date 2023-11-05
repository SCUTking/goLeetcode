package random

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	listNodes := make([]*ListNode, 0)
	if head == nil {
		return nil
	}
	for head.Next != nil {
		for _, node := range listNodes {
			if node == head {
				return node
			}
		}
		listNodes = append(listNodes, head)
		head = head.Next
	}

	return nil

}
