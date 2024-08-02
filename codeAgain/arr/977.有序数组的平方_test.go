package arr

import (
	"reflect"
	"testing"
)

func Test_sortedArrSquare(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			name: "",
			args: args{
				nums: []int{-2, -2, -1, 0, 2, 3},
			},
			want: []int{0, 1, 4, 4, 4, 9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortedArrSquare(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortedArrSquare() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_best_sortedArrSquare(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			name: "",
			args: args{
				nums: []int{-2, -2, -1, 0, 2, 3},
			},
			want: []int{0, 1, 4, 4, 4, 9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := best_sortedArrSquare(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("best_sortedArrSquare() = %v, want %v", got, tt.want)
			}
		})
	}
}
