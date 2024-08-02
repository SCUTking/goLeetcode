package random2

import (
	"fmt"
	"strings"
	"testing"
)

func Test(t *testing.T) {
	temp := make([][]int, 0)
	temp = append(temp, []int{0, 1})
	temp = append(temp, []int{1, 0})
	//temp = append(temp, []int{1, 3})
	//temp = append(temp, []int{2, 0})
	//temp = append(temp, []int{2, 3})
	transportationHub(temp)
}

func Test1(t *testing.T) {
	//predictPartyVictory("DDRRR")
	a := "00001"
	prefix := strings.TrimLeft(a, "0")
	fmt.Println(prefix)
}
func Test2(t *testing.T) {
	addBinary("1010", "1011")
}
func Test3(t *testing.T) {
	var str string = "abs"
	s := str[:2]
	fmt.Println(str[:2] == "ab")
	fmt.Println(s)
}

func abs(a int32) int32 {
	//判断int32的首位是1还是0
	fmt.Println((-4) ^ (-2))
	return 1
}

type Inter1 interface {
	fun1()
	fun2()
}

type struct1 struct {
	Inter1
}

func (*struct1) fun1() {
	fmt.Println("1")
}

func Test5(t *testing.T) {
	var s struct1
	//s.fun1()
	s.fun2()
}
