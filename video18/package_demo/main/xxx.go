package main

import (
	nazha "code/video18/package_demo/calc" // 给导入的包起别名
	"fmt"
)

// Go语言中不允许导入包而不使用！！！
// Go语言中不允许循环引用包！！！

//// 当你的代码都放在$GOPATH目录下的话
//// 我们导入包的路径要从$GOPATH/src后面继续写起
//import "code/video18/package_demo/calc"

// 多用来做一些初始化的操作
func init() {
	fmt.Println("main.init()")
}

func main() {
	fmt.Println("hello")
	ret := nazha.Add(10, 20)
	fmt.Println(ret)
	fmt.Println(nazha.Name)
}
