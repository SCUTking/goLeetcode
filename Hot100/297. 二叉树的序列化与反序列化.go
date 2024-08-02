package Hot100

import (
	"strconv"
	"strings"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Codec struct {
}

func Constructor2() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	var reslut strings.Builder
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		if cur == nil {
			reslut.WriteString("null")
			reslut.WriteString(",")
			continue
		}
		reslut.WriteString(strconv.Itoa(cur.Val))
		reslut.WriteString(",")

		queue = append(queue, cur.Left)
		queue = append(queue, cur.Right)
	}
	s := reslut.String()
	s = s[:len(s)-1]
	return s
}

// 层序遍历构造二叉树
// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	split := strings.Split(data, ",")
	if len(split) == 0 {
		return nil
	}
	atoi, _ := strconv.Atoi(split[0])
	root := &TreeNode{
		Val: atoi,
	}
	queue := make([]*TreeNode, 0)
	index := 1
	queue = append(queue, root)
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		if split[index] != "null" {
			i, _ := strconv.Atoi(split[index])
			cur.Left = &TreeNode{
				Val: i,
			}
			queue = append(queue, cur.Left)
		}
		index++
		if split[index] != "null" {
			i, _ := strconv.Atoi(split[index])
			cur.Right = &TreeNode{
				Val: i,
			}
			queue = append(queue, cur.Right)
		}
		index++

	}

	return root
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
