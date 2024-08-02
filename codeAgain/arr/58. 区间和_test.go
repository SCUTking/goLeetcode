package arr

import "testing"

func Test_sumInArea(t *testing.T) {
	type args struct {
		nums []int
		s    int
		e    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "",
			args: args{nums: []int{1, 2, 3, 4, 5}, s: 1, e: 3},
			want: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumInArea(tt.args.nums, tt.args.s, tt.args.e); got != tt.want {
				t.Errorf("sumInArea() = %v, want %v", got, tt.want)
			}
		})
	}
}
