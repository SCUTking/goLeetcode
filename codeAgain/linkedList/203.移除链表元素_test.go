package linkedList

import (
	"reflect"
	"testing"
)

func Test_removeElement(t *testing.T) {
	type args struct {
		head   *ListNode
		target int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		// TODO: Add test cases.
		{
			name: "",
			args: args{head: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}}, target: 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeElement(tt.args.head, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeElement() = %v, want %v", got, tt.want)
			}
		})
	}
}
