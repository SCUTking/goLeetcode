package random1

import (
	"strconv"
	"strings"
)

func DecodeString(s string) string {
	stack := make([]string, 0)
	for i := 0; i < len(s); i++ {
		each := s[i]
		if each != ']' {
			stack = append(stack, string(each))
		} else {
			innerString := ""
			innerArr := make([]string, 0)
			for stack[len(stack)-1] != "[" {
				innerArr = append(innerArr, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			for i := len(innerArr) - 1; i >= 0; i-- {
				innerString += innerArr[i]
			}
			//去除最后的【
			stack = stack[:len(stack)-1]
			numString := ""
			for len(stack) > 0 && stack[len(stack)-1] >= "0" && stack[len(stack)-1] <= "9" {
				numString += stack[len(stack)-1]
				stack = stack[:len(stack)-1]
			}
			num := ReverseString(numString)
			atoi, _ := strconv.Atoi(num)

			if atoi > 1 {
				temp := innerString
				for i := 1; i < atoi; i++ {
					innerString += temp
				}
			}
			stack = append(stack, innerString)

			strings.Repeat("1", 1)
		}
	}

	ans := ""
	for _, s2 := range stack {
		ans += s2
	}
	return ans

}
func ReverseString(s string) string {
	runes := []rune(s)
	length := len(runes)
	for i := 0; i < length/2; i++ {
		runes[i], runes[length-1-i] = runes[length-1-i], runes[i]
	}

	return string(runes)
}
