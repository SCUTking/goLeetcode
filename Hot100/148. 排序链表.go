package Hot100

import "sort"

func sortList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	s := make([]*ListNode, 0)
	for head != nil {
		s = append(s, head)
		head = head.Next
	}

	sort.Slice(s, func(i, j int) bool {
		return s[i].Val < s[j].Val
	})

	for i := 0; i < len(s); i++ {
		s[i].Next = nil
	}

	for i := 0; i < len(s)-1; i++ {
		if i != len(s)-1 {
			s[i].Next = s[i+1]
		} else {
			s[i].Next = nil
		}
	}

	return s[0]

}
