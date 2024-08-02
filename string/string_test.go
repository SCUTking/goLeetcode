package string

import (
	"fmt"
	"strings"
	"testing"
)

func Test678(t *testing.T) {
	CheckValidString("(((((*)))**")
	var a1 string
	fmt.Print(a1)
	var a strings.Builder
	a.WriteString("asd") //在后面拼接上asd
	a.Grow(10)           //扩容10个byte数组大小
	a.Reset()            //重置builder

}
