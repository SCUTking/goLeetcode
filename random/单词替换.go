package random

import "strings"

func ReplaceWords(dictionary []string, sentence string) string {
	type node struct {
		next  [26]*node
		isEnd int
	}
	root := &node{}
	temp := root
	for _, s := range dictionary {
		root = temp
		for _, i := range s {
			i2 := i - 'a'
			if root.next[i2] == nil {
				root.next[i2] = &node{}
			}
			root = root.next[i2]
		}
		root.isEnd++
	}

	split := strings.Split(sentence, " ")

	var findIsEnd func(s string, root *node) string
	findIsEnd = func(s string, root *node) string {
		res := []rune{}
		for _, i := range s {
			i2 := i - 'a'
			if root.next[i2] != nil {
				res = append(res, i)
				if root.next[i2].isEnd > 0 {

					return string(res)
				}
				root = root.next[i2]
			} else {
				return s
			}
		}
		return s

	}
	ans := []string{}
	for _, s := range split {
		end := findIsEnd(s, temp)
		ans = append(ans, end)
	}

	return strings.Join(ans, " ")
}
