package dynamic_programming

import (
	"fmt"
	"testing"
)

func TestDP(t *testing.T) {
	change(4, []int{1, 2, 3})
	fmt.Println(".........................")
	combinationSum4([]int{1, 2, 3}, 4)

	var arr [2]int
	arr[0] = 1
	arr[1] = 2

	s := arr[:1]
	fmt.Println(s)

}

func TestP(t *testing.T) {
	wordBreak("catsandog", []string{"cats", "dog", "sand", "and", "cat"})
}
