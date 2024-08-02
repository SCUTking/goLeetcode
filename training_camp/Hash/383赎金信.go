package Hash

func canConstruct(ransomNote string, magazine string) bool {
	m := make(map[int32]int)
	for _, i := range magazine {
		m[i-'a']++
	}
	for _, i := range ransomNote {
		m[i-'a']--
		if i < 0 {
			return false
		}
	}
	return true
}
