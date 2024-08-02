package arr

import "testing"

func Test_removeTarget(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "",
			args: args{
				nums:   []int{0, 1, 2, 2, 3, 0, 4, 2},
				target: 2,
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeTarget(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("removeTarget() = %v, want %v", got, tt.want)
			}
		})
	}
}
