package hash

import (
	"goLeetcode/skills"
	"strconv"
	"strings"
)

func CountDistinctIntegers(nums []int) int {
	set := make(map[int]string, len(nums))
	for _, v := range nums {
		set[v] = ""
		rev := skills.Rev(v)
		set[rev] = ""
	}
	return len(set)
}

func fanzhuan(num int) int {
	itoa := strconv.Itoa(num)

	s := reverseString(itoa)

	atoi, _ := strconv.Atoi(s)
	return atoi
}

func reverseString(s string) string {
	runes := []rune(s)
	length := len(runes)
	for i := 0; i < length/2; i++ {
		runes[i], runes[length-1-i] = runes[length-1-i], runes[i]
	}

	str := string(runes)

	return strings.TrimLeft(str, "0")
}
