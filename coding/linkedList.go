package coding

//1.删除链表的倒数第 n 个结点
//给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。
//示例 1：
//输入：head = [1,2,3,4,5], n = 2
//输出：[1,2,3,5]
//示例 2：
//输入：head = [1], n = 1
//输出：[]
//示例 3：
//输入：head = [1,2], n = 1
//输出：[1]

type LinkedNode struct {
	Val  int
	Next *LinkedNode
}

func delete(head *LinkedNode, n int) {
	l := 0
	temp := head
	for temp != nil {
		l++
		temp = temp.Next
	}

	count := l - n
	pre := &LinkedNode{Next: head}
	temp = head
	for count > 0 {
		count--
		pre = pre.Next
		temp = temp.Next
	}

	pre.Next = temp.Next

}
