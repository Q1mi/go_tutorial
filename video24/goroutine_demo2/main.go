package main

import (
	"fmt"
	"sync"
)

// goroutine demo2

var wg sync.WaitGroup

func main() { // 开启一个主goroutine去执行main函数

	wg.Add(10000) // 计数牌+1
	for i := 0; i < 10000; i++ {
		go func(i int) {
			fmt.Println("hello", i)
			wg.Done()
		}(i)
	}
	fmt.Println("hello main")
	wg.Wait() // 阻塞，等所有小弟都干完活之后才收兵
}
