package fileOperation

import (
	"fmt"
	"io"
	"os"
	"testing"
)

type Shape interface {
	Area() float64
}
type Shape1 interface {
	Area() float64
	test() int
}
type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) test() int {
	return 1
}

func main() {

}
func Test1(t *testing.T) {
	var w io.Writer
	w = os.Stdout
	w.Write([]byte{0, 1})
	fmt.Println(w)
	rw, ok := w.(io.ReadWriter) // success: *os.File has both Read and Write
	rw.Read([]byte{0, 1})
	fmt.Println(rw)
	fmt.Println(ok)

}

func Test2(t *testing.T) {
	var shape Shape
	shape = Rectangle{Width: 5, Height: 10}
	var f func(shape1 Shape1)
	f = func(shape1 Shape1) {
		shape1.test()
	}
	if in, ok := shape.(Shape1); ok {
		f(in)
		fmt.Println("shape is a Shape interface", in)
	} else {
		fmt.Println("shape is not a Shape interface")
	}

}

func TestByLine(t *testing.T) {
	lineByLine1("./test.txt")
}

func TestByWord(t *testing.T) {
	lineByWord("./test.txt")
}

func TestByChar(t *testing.T) {
	lineByChar("./test.txt")
}

func TestAppendWrite(t *testing.T) {
	file, err := os.OpenFile("./test1.txt", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		return
	}
	defer file.Close()
	_, err = file.Write([]byte("1.1.1.1\n"))
	_, err = file.Write([]byte("1.1.1.1\n"))
	if err != nil {
		fmt.Println(err)
	}
}
