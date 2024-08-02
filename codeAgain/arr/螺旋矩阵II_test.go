package arr

import (
	"reflect"
	"testing"
)

func Test_matrix(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		// TODO: Add test cases.
		{
			name: "",
			args: args{n: 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := matrix(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("matrix() = %v, want %v", got, tt.want)
			}
		})
	}
}
