package sort

import "testing"

func TestHeapSortMulti(t *testing.T) {
	type args struct {
		nums []student
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "",
			args: args{nums: []student{
				{
					age: 10,
				},
				{
					age: 100,
				},
				{
					age: 199,
				},
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			HeapSortMulti(tt.args.nums)
		})
	}
}
