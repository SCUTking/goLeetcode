package String

func reverseStr(s string, k int) string {
	bytes := []byte(s)
	n := len(s)
	for i := 2*k - 1; i < n; i += 2 * k {
		reverse(bytes, i-2*k+1, i-k)
	}
	modNum := n % (2 * k)
	if modNum >= k {
		reverse(bytes, n-modNum, n-modNum+k-1)
	} else if modNum == 0 {
		return string(bytes)
	} else {
		reverse(bytes, n-modNum, n-1)
	}

	return string(bytes)
}

func reverse(s []byte, start, end int) {
	n := (end - start + 1)
	for i := 0; i < n/2; i++ {
		s[start+i], s[end-i] = s[end-i], s[start+i]
	}
}
func reverseStr1(s string, k int) string {
	t := []byte(s)
	for i := 0; i < len(s); i += 2 * k {
		sub := t[i:min(i+k, len(s))] //浅拷贝——只拷贝地址，可以修改实际的
		for j, n := 0, len(sub); j < n/2; j++ {
			sub[j], sub[n-1-j] = sub[j], sub[n-1-j]
		}
	}
	return string(t)
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a

}
