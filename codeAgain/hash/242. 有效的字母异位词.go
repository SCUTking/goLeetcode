package hash

// 算法要点：最后判断条件为查看hash表的长度，所以当hash的value为0后得删除对应的表
// 时间复杂度：O（n）
// 空间复杂度：O（n）
// 未能考虑到的点：可以先用长度是否相等进行优化
func isAnagram(s string, t string) bool {
	m := make(map[uint8]int, 0)
	for i := 0; i < len(s); i++ {
		m[s[i]]++
	}

	for i := 0; i < len(t); i++ {
		m[t[i]]--
		if m[t[i]] < 0 {
			return false
		} else if m[t[i]] == 0 {
			delete(m, t[i])
		}
	}
	return len(m) == 0
}
