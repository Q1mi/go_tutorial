package main

import (
	"fmt"
	"strings"
)

//匿名函数和闭包
//定义一个函数它的返回值是一个函数
//把函数作为返回值
func a(name string) func() {
	return func() {
		fmt.Println("hello", name)
	}
}

func makeSuffixFunc(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

func calc(base int) (func(int) int, func(int) int) {
	add := func(i int) int {
		base += i
		return base
	}

	sub := func(i int) int {
		base -= i
		return base
	}
	return add, sub
}

func main() {
	//闭包 = 函数 + 外层变量的引用
	x, y := calc(100)
	ret1 := x(200) //base = 100 + 200
	fmt.Println(ret1)
	ret2 := y(200) //base = 300 - 200
	fmt.Println(ret2)
}
