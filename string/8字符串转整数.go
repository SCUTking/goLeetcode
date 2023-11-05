package string

import (
	"math"
	"strconv"
	"strings"
)

func MyAtoi(s string) int {

	s = strings.TrimLeft(s, "0")
	s = strings.TrimLeft(s, " ")
	runes := []rune(s)
	for i, _ := range runes {
		if runes[i] < '0' || runes[i] > '9' || (runes[i] == '+' && i != 0) || (runes[i] == '-' && i != 0) {
			if runes[i] == '+' || runes[i] == '-' {
				continue
			}
			if runes[i] == ' ' {
				runes = runes[:i]
				break
			}
			return 0

		}
	}
	atoi, _ := strconv.Atoi(string(runes))
	if atoi < math.MinInt32 {
		return math.MinInt32
	}
	if atoi > math.MaxInt32 {
		return math.MaxInt32
	}

	return atoi

}
