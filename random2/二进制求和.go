package random2

import (
	"fmt"
	"strings"
)

func addBinary(a string, b string) string {
	charA := []rune(ReverseString(a))
	charB := []rune(ReverseString(b))

	isOver := false //是否进位
	res := make([]rune, Max(len(a), len(b))+1)
	for i := 0; i < Max(len(a), len(b))+1; i++ {
		sum := 0
		if isOver {
			sum += 1
		}
		if i < len(b) && charB[i] == '1' {
			sum += 1
		}
		if i < len(a) && charA[i] == '1' {
			sum += 1
		}

		if sum >= 2 {
			isOver = true
			res[i] = int32(sum-2) + '0'
		} else {
			isOver = false
			res[i] = int32(sum) + '0'
		}
	}

	s := string(res)
	fmt.Println(s)
	return OutPrefix(ReverseString(string(res)))

}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func OutPrefix(a string) string {
	prefix := strings.TrimLeft(a, "0")
	return prefix
}

func Max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func ReverseString(a string) string {

	runes := []rune(a)
	for i := 0; i < len(runes)/2; i++ {
		runes[i], runes[len(runes)-i-1] = runes[len(runes)-i-1], runes[i]
	}

	return string(runes)
}
