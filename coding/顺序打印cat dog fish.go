package coding

import (
	"fmt"
	"sync"
	"time"
)

func print1() {
	c := make([]chan int, 3)
	for i := 0; i < 3; i++ {
		c[i] = make(chan int, 1)
	}

	var wg sync.WaitGroup
	wg.Add(3)
	for i := 1; i <= 3; i++ {
		go func(i int) {
			defer wg.Done()
			for true {
				<-c[i-1]
				fmt.Printf("%d", i)
				c[i%3] <- 1
			}
		}(i)
	}

	c[0] <- 1
	time.Sleep(time.Second)
}
