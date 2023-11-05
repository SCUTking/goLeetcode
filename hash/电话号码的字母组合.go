package hash

import (
	"strconv"
)

func LetterCombinations(digits string) []string {

	result := make([]string, 0, 10)
	mapletter := make(map[int32][]string, 8)
	mapletter[2] = []string{"a", "b", "c"}
	mapletter[3] = []string{"d", "e", "f"}
	mapletter[4] = []string{"g", "h", "i"}
	mapletter[5] = []string{"j", "k", "l"}
	mapletter[6] = []string{"m", "n", "o"}
	mapletter[7] = []string{"p", "q", "r"}
	mapletter[8] = []string{"s", "t", "u"}
	mapletter[9] = []string{"v", "w", "y", "z"}

	if len(digits) == 0 {
		return []string{}
	}
	back(digits, 0, "", mapletter, &result)
	return result
}

func back(digits string, index int, Combination string, mapA map[int32][]string, result *[]string) {
	if index == len(digits) {
		*result = append(*result, Combination)
	} else {
		digit, _ := strconv.Atoi(string(digits[index]))

		strings := mapA[int32(digit)]
		length := len(strings)
		for i := 0; i < length; i++ {
			back(digits, index+1, Combination+strings[i], mapA, result)
		}
	}
}
