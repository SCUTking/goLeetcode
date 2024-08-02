package Hot100

import (
	"fmt"
	"math"
	"reflect"
	"testing"
)

func Test15(t *testing.T) {
	arr := []int{0, 1}
	var test func(num []int)
	test = func(num []int) {
		num[0] = 1
	}

	a, _ := a1()
	fmt.Println(&a)
	a, c := a1()
	fmt.Println(&a, c)
	test(arr)
	println(arr[0])
	//threeSum([]int{-1, 0, 1, 2, -1, -4})
}

func a1() (int, int) {
	return 1, 1
}
func Test1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("nihaoya", r)
		}
	}()

	fmt.Println("1")
	panic("nihao")
	fmt.Println(2)
}

type MyStack struct {
	inQueue, outQueue []int
}

func Test2(t *testing.T) {
	var ints []int
	ints = append(ints, 1)
	fmt.Println(ints)
}

func Test111(t *testing.T) {
	//treeNode := TreeNode{
	//	Val:   1,
	//	Left:  nil,
	//	Right: nil,
	//}

	constructor2 := Constructor2()
	serialize := constructor2.serialize(nil)
	fmt.Println(serialize)

	deserialize := constructor2.deserialize(string(serialize))
	fmt.Println(deserialize)
}

func Test3(t *testing.T) {
	m := make(map[int]int, 0)
	i := m[0]
	fmt.Println(i)
}
func Test4(t *testing.T) {
	chan1 := make(chan int)
	chan2 := make(chan int)
	close(chan1)
	close(chan2)
	select {
	case <-chan1:
		fmt.Println("chan1")
	case chan2 <- 1:
		fmt.Println("chan1")
	default:
		fmt.Println("default")
	}
}
func Test5(t *testing.T) {
	var chan1 chan int
	chan1 <- 1
}

type struct1 struct {
	m int
}

func Test7(t *testing.T) {
	s := struct1{m: 1}
	fmt.Printf("s=%+v", s)
	fmt.Printf("s=%#v", s)
	fmt.Printf("s=%v", s)
	s1 := struct1{m: 1}
	fmt.Printf("s=%v", s1)
	ints := [1]int{1}
	ints1 := [1]int{1}
	fmt.Println(ints == ints1)
}

func Test6(t *testing.T) {
	test := make([]int, 0)
	test = append(test, 1)
	//for i := 0; i < len(test); i++ {
	//	test = append(test, 2)
	//	fmt.Println(test[i])
	//}
	for _, i := range test {
		test = append(test, 2)
		fmt.Println(i)
	}

}
func Test78(t *testing.T) {
	test := make([]int, 0)
	test = append(test, 1)
	//for i := 0; i < len(test); i++ {
	//	test = append(test, 2)
	//	fmt.Println(test[i])
	//}
	for _, i := range test {
		test = append(test, 2)
		fmt.Println(i)
	}
}
func Test781(t *testing.T) {
	var a func(a int) int
	var b func(a int) int

	fmt.Println(reflect.DeepEqual(a, b))
}
func Test72(t *testing.T) {
	m := make(map[int]int, 0)
	i := m[0]
	fmt.Println(i)
}
func Test722(t *testing.T) {
	m := make(map[int]int, 0)
	go func() {
		for true {
			m[1] = 1
		}
	}()
	go func() {
		for true {
			i := m[1]
			fmt.Println(i)
		}
	}()
	select {}
}
func Test723(t *testing.T) {
	m := make(map[int]int, 0)
	//增

	//删
	delete(m, 0)
	//改

	//查
	i, ok := m[0]
	if !ok {
		fmt.Println(i)
	}
}
func Test724(t *testing.T) {
	var m map[int]int
	//增
	//m[0] = 1

	i := m[1]
	fmt.Println(i)
	//删
	//delete(m, 0)
	//改

	//查
	//i, ok := m[0]
	//if !ok {
	//	fmt.Println(i)
	//}
}
func Test725(t *testing.T) {
	mm := make(map[string]string, 0)
	s := mm["1"]
	fmt.Printf("%v", s)
}
func Test625(t *testing.T) {

	fmt.Println(math.Pow10(1))
}
func Test36(t *testing.T) {
	// 初始化无向图
	graph := [][]int{
		{-1, 2, -1, 6, -1},
		{2, -1, 3, 8, 5},
		{-1, 3, -1, -1, 7},
		{6, 8, -1, -1, 9},
		{-1, 5, 7, 9, -1},
	}

	startNode := 0 // 起始节点为 0
	distances := dijkstra(graph, startNode)

	// 输出起始节点到其他节点的最短距离
	fmt.Println("Shortest distances from node", startNode, "to all other nodes:")
	for i, d := range distances {
		fmt.Printf("Node %d: %d\n", i, d)
	}

}

type MyStack1 struct {
	inQueue, outQueue []int
}

func Constructor1() MyStack1 {
	return MyStack1{}
}

func Test21(t *testing.T) {
	constructor := Constructor1()
	var temp []int
	fmt.Println(temp == nil)
	fmt.Println(len(constructor.inQueue))

}
func Test22(t *testing.T) {
	//var f func(a [1000]int64)
	//f = func(a [1000]int64) {
	//	f(a)
	//}
	//f([1000]int64{})

	var f func()
	go f()

}
