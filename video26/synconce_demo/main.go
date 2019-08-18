package main

import (
	"fmt"
	"sync"
)

// sync.Once 在并发场景下确保某个函数只执行一次

var wg sync.WaitGroup
var once sync.Once

func f1(ch1 chan<- int) {
	defer wg.Done()
	for i := 0; i < 100; i++ {
		ch1 <- i
	}
	close(ch1)
}

func f2(ch1 <-chan int, ch2 chan<- int) {
	defer wg.Done()
	for {
		x, ok := <-ch1
		if !ok {
			break
		}
		ch2 <- x * x
	}
	// 确保在并发的场景下 close(ch2)这个操作只执行一次
	f := func() {
		close(ch2) // 闭包
	}
	once.Do(f)
	//close(ch2)
}

func main() {
	a:= make(chan int, 100)
	b := make(chan int, 100)
	wg.Add(3)
	go f1(a)
	go f2(a, b)
	go f2(a, b)
	wg.Wait()
	for ret := range b{
		fmt.Println(ret)
	}
}
