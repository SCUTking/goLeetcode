package String

func strStr1(haystack string, needle string) int {
	const base = 1000000
	m, n := len(haystack), len(needle)
	if n > m {
		return -1
	}
	if n == 0 {
		return 0
	}

	//26进制数最高位的乘数
	power := 1
	for i := 1; i < n; i++ {
		power = (power * 26) % base
	}
	targetCode := hash(needle)
	firstCode := hash(haystack[0:n])
	if targetCode == firstCode {
		return 0
	}
	for i := 1; i <= m-n; i++ {
		firstCode = firstCode - (int(haystack[i-1]-'a')*power)%base
		if firstCode < 0 {
			firstCode += base
		}
		firstCode = (firstCode*26 + int(haystack[i+n-1]-'a')) % base
		//如果hash值相等，就使用双重检查，避免hash冲突
		k := 0
		if firstCode == targetCode {
			for ; k < n; k++ {
				if haystack[i+k] == needle[k] {
					continue
				} else {
					break
				}
			}
		}
		if k == n {
			return i
		}

	}

	return -1
}
func hash(s string) int {
	res := 0
	base := 1000000
	//res*26相当于将26进制的hash值每一位左移一位
	//
	for _, i2 := range s {
		i3 := i2 - 'a'
		res = (res*26 + int(i3)) % base
	}
	return res
}
