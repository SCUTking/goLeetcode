package coding

import (
	"fmt"
	"sync"
)

func print3() {
	var n int = 5                  //GoRoutine的个数
	chanArr := make([]chan int, n) //每个协程一个管道
	//初始化管道数组
	for i := 0; i < n; i++ {
		chanArr[i] = make(chan int, 1)
	}

	//输出的最大最小值
	var minNum, maxNum int
	minNum = 1
	maxNum = 100

	var wg sync.WaitGroup
	wg.Add(n)

	for i := 0; i < n; i++ {
		go func(i int) {
			defer wg.Done()
			defer close(chanArr[(i+1)%n])
			//循环，直到被管道上一个线程关闭
			for num := range chanArr[i] {
				if num > maxNum {
					break
				}
				fmt.Printf("g %v 打印的数字%d\n", i, num)
				chanArr[(i+1)%n] <- num + 1
			}

		}(i)
	}
	//main 协程进行开始
	chanArr[0] <- minNum

	wg.Wait()
}
