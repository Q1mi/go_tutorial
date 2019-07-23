package main

import "fmt"

// 指针

/*
func main() {
	a := 10 // int类型

	b := &a                      // *int类型（int指针）
	fmt.Printf("%v %p\n", a, &a) // 10 0xc00001a098
	fmt.Printf("%v %p\n", b, b)  // 0xc00001a098 0xc00001a098
	// 变量b是一个int类型的指针（*int），它存储的是变量a的内存地址
	c := *b        // 根据内存地址去取值
	fmt.Println(c) // 10
}
*/

func modify1(x int) {
	x = 100
}

func modify2(y *int) {
	*y = 100
}

func main() {
	a := 1
	modify1(a)
	fmt.Println(a) // 1
	modify2(&a)
	fmt.Println(a)
}
