package main

import "goLeetcode/linkedList"

func main() {

	l := &linkedList.ListNode{Val: 2, Next: &linkedList.ListNode{Val: 1}}
	l1 := &linkedList.ListNode{Val: 3, Next: l}
	//l2 := &linkedList.ListNode{Val: 3}
	l3 := &linkedList.ListNode{Val: 4, Next: l1}

	//linkedList.AddTwoNumbers(l3, l1)
	//linkedList.SwapPairs(l3)
	linkedList.RotateRight(l3, 1)

}
