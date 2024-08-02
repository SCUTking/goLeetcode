package random2

func longestPrefix(s string) string {
	n := len(s)
	res := ""
	for i := 0; i < n-1; i++ {
		prefix := s[:i+1]
		suffix := s[n-i-1:]
		if prefix == suffix {
			res = prefix
		}
	}

	return res
}
