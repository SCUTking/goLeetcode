package training_camp

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

// 每行数字数量固定
// 不知道行数
func Test1(t *testing.T) {
	a := 0
	b := 0
	for {
		n, _ := fmt.Scan(&a, &b)
		fmt.Scan(&a, &b)
		if n == 0 {
			break
		} else {
			fmt.Printf("%d\n", a+b)
		}
	}
}
func Test2(t *testing.T) {

	v := new([]int)
	fmt.Println(*v) //
	fmt.Println(v)  //0xc00004c088

}

func Benchmark1(b *testing.B) {
	n := b.N
	fmt.Println(n)
	for i := 0; i < n; i++ {
		fmt.Println("非耗时工作")
	}
	assert.Equal(b, 1, 1, "比对成功")
}
