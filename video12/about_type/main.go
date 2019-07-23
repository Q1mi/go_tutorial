package main

import "fmt"

// 自定义类型和类型别名示例

// 1. 自定义类型

// MyInt 基于int类型的自定义类型
type MyInt int

// 2. 类型别名

// NewInt int类型别名
type NewInt = int

func main() {
	var i MyInt
	fmt.Printf("type:%T value:%v\n", i, i)

	var x NewInt
	fmt.Printf("type:%T value:%v\n", x, x)
}
