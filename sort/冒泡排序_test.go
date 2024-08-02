package sort

import "testing"

func TestBubbleSortOne(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "",
			args: args{nums: []int{2, 1, 3, 0}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			BubbleSortOne(tt.args.nums)
		})
	}
}
