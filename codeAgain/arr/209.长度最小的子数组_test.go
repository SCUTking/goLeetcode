package arr

import "testing"

func TestSmallestSubArr(t *testing.T) {
	type args struct {
		nums []int
		s    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "",
			args: args{nums: []int{2, 3, 1, 2, 4, 3}, s: 7},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SmallestSubArr(tt.args.nums, tt.args.s); got != tt.want {
				t.Errorf("SmallestSubArr() = %v, want %v", got, tt.want)
			}
		})
	}
}
