package arr

import "testing"

func Test_divisionLand(t *testing.T) {
	type args struct {
		land [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "",
			args: args{land: [][]int{[]int{1, 2, 3}, []int{2, 1, 3}, []int{1, 2, 3}}},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := divisionLand(tt.args.land); got != tt.want {
				t.Errorf("divisionLand() = %v, want %v", got, tt.want)
			}
		})
	}
}
