package MultiThread

import "fmt"

type FooBar struct {
	n int
}

var c []chan int = make([]chan int, 2)

func NewFooBar(n int) *FooBar {
	c[0] = make(chan int, 1)
	c[1] = make(chan int, 1)
	c[0] <- 1
	return &FooBar{n: n}
}

func (fb *FooBar) Foo() {
	for i := 0; i < fb.n; i++ {
		// printFoo() outputs "foo". Do not change or remove this line.
		printFoo()
	}
}

func printFoo() {
	<-c[0]
	fmt.Printf("%v", "foo")
	c[1] <- 1
}

func (fb *FooBar) Bar() {
	for i := 0; i < fb.n; i++ {
		// printBar() outputs "bar". Do not change or remove this line.
		printBar()
	}
}

func printBar() {
	<-c[1]
	fmt.Printf("%v\n", "bar")
	c[0] <- 1
}
