package string

func CheckValidString1(s string) bool {

	var leftStk, asteriskStk []int
	for i, ch := range s {
		if ch == '(' {
			leftStk = append(leftStk, i)
		} else if ch == '*' {
			asteriskStk = append(asteriskStk, i)
		} else if len(leftStk) > 0 {
			leftStk = leftStk[:len(leftStk)-1]
		} else if len(asteriskStk) > 0 {
			asteriskStk = asteriskStk[:len(asteriskStk)-1]
		} else {
			return false
		}

	}
	i := len(leftStk) - 1
	for j := len(asteriskStk) - 1; j >= 0 && i >= 0; i, j = i-1, j-1 {
		if leftStk[i] > asteriskStk[j] {
			return false
		}
	}
	return i < 0
}
