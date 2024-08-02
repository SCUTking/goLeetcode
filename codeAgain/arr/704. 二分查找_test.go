package arr

import "testing"

func Test_sortSearch(t *testing.T) {
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
				nums:   []int{0, 2, 3, 5, 6, 8, 9},
				target: 9,
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortSearch(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("sortSearch() = %v, want %v", got, tt.want)
			}
		})
	}
}
