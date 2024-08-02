package sort

import (
	"fmt"
	"testing"
)

func Test_selectionSort(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{name: "1", args: args{[]int{2, 2, 1, 3, 1, 5}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			selectionSort(tt.args.nums)
			fmt.Println(tt.args.nums)
		})
	}
}
