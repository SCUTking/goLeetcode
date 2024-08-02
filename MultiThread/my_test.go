package MultiThread

import (
	"testing"
	"time"
)

func Test1115(t *testing.T) {

	bar := FooBar{n: 10}
	go func() {
		bar.Bar()
	}()
	go func() {
		bar.Foo()
	}()

	time.Sleep(5 * time.Second)

}
