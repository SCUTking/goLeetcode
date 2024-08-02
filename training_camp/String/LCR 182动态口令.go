package String

func dynamicPassword(password string, target int) string {
	//多次翻转
	bytes := []byte(password)
	reverse1(bytes, 0, target-1)
	reverse1(bytes, target, len(password)-1)
	reverse1(bytes, 0, len(password)-1)
	return string(bytes)
}
func reverse1(s []byte, start, end int) {
	n := end - start + 1
	for i := 0; i < n/2; i++ {
		s[start+i], s[end-i] = s[end-i], s[start+i]
	}
}
