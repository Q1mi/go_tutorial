package main

import "fmt"

// 为什么需要接口

type dog struct{}

func (d dog) say() {
	fmt.Println("汪汪汪~")
}

type cat struct{}

func (c cat) say() {
	fmt.Println("喵喵喵~")
}

type person struct {
	name string
}

func (p person) say() {
	fmt.Println("啊啊啊~")
}

// 接口不管你是什么类型，它只管你要实现什么方法
// 定义一个类型，一个抽象的类型，只要实现了say()这个方法的类型都可以称为sayer类型
type sayer interface {
	say()
}

// 打的函数
func da(arg sayer) {
	arg.say() // 不管是传进来的是什么，我都要打Ta，打Ta Ta就会叫，就要执行Ta的say方法
}

func main() {
	//c1 := cat{}
	//da(c1)
	//d1 := dog{}
	//da(d1)
	//p1 := person{
	//	name: "娜扎",
	//}
	//da(p1)

	var s sayer
	c2 := cat{}
	s = c2
	p2 := person{name: "小王子"}
	s = p2
	fmt.Println(s)
}
