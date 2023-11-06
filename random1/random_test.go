package random1

import (
	"fmt"
	"testing"
)

func Test324(f *testing.T) {
	WiggleSort([]int{1, 2, 3})
}

func TestLCR163(f *testing.T) {
	FindKthNumber(190)
}

func Test143(f *testing.T) {
	ReorderList(&ListNode{Val: 2, Next: &ListNode{3, &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: nil}}}})
}
func Test394(f *testing.T) {
	DecodeString("3[z]2[2[y]pq4[2[jk]e1[f]]]ef")
}

func Test73(f *testing.T) {
	Subsets([]int{1, 2, 3})
}

func Test(f *testing.T) {
	//num := 1
	//p := &num
	//fmt.Printf("大小为%v", unsafe.Sizeof(p))
	//
	//type test struct {
	//	age int
	//}
	//
	//t := test{age: 1}
	//a := new(test)
	//a = &t
	//a.age = 100
	//
	//fmt.Print(t)

	ints := make([]int, 10)
	fmt.Printf("len %d,cap %d", len(ints), cap(ints))

}
