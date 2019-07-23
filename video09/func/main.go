package main

import "fmt"

//函数
//定义一个不需要参数也没有返回值的函数：sayHello
func sayHello() {
	fmt.Println("Hello 沙河小王子！")
}

//定义一个接收string类型的name参数
func sayHello2(name string) {
	fmt.Println("Hello ", name)
}

//定义接收多个参数的函数并且有一个返回值
func intSum(a int, b int) int {
	ret := a + b
	return ret
}

func intSum2(a int, b int) (ret int) {
	ret = a + b
	return
}

//函数接收可变参数,在参数名后面加... 表示可变参数
// 可变参数在函数体中是切片类型
func intSum3(a ...int) int {
	ret := 0
	for _, arg := range a {
		ret = ret + arg
	}
	return ret
}

// 固定参数和可变参数同时出现时，可变参数要放在最后
// go语言的函数中没有默认参数
func intSum4(a int, b ...int) int {
	ret := a
	for _, arg := range b {
		ret = ret + arg
	}
	return ret
}

//Go语言中函数参数类型简写
func intSum5(a, b int) (ret int) {
	ret = a + b
	return
}

//定义具有多个返回值的函数,多返回值也支持类型简写。
func calc(a, b int) (sum, sub int) {
	sum = a + b
	sub = a - b
	return
}

func main() {
	//函数调用
	x, y := calc(100, 200)
	fmt.Println(x, y)
}
