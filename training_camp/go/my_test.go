package _go

import (
	"fmt"
	"sync"
	"testing"
)

func Test1(t *testing.T) {
	var m1 sync.Mutex
	var m2 sync.Mutex
	m2.Lock()
	count := -1

	go func() {
		for true {
			m1.Lock()
			count++
			fmt.Println("G1:", count)
			m2.Unlock()
		}
	}()
	go func() {
		for true {
			m2.Lock()
			count++
			fmt.Println("G2:", count)
			m1.Unlock()
		}
	}()
	select {}
}
