package coding

import (
	"fmt"
	"sync"
)

func print2() {
	strChan := make(chan int, 1)
	numChan := make(chan int, 1)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 65; i < 65+26; i++ {
			<-strChan
			fmt.Printf("%v ", string(byte(i)))
			numChan <- 1
		}
	}()

	//main的gproutinue发布号令
	strChan <- 1

	go func() {
		defer wg.Done()
		for i := 1; i <= 26; i++ {
			<-numChan
			fmt.Printf("%v ", i)
			strChan <- 1
		}
	}()

	wg.Wait()
}
