package sort

import "testing"

func TestMS(t *testing.T) {
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
			args: args{[]int{8, 4, 5, 7, 1, 3, 6, 2}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			MS(tt.args.nums)
		})
	}
}
