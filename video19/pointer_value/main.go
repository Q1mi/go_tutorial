package main

import "fmt"

// 使用值接收者实现接口和使用指针接收者实现接口的区别

// 接口的嵌套
type animal interface {
	mover
	sayer
}

type mover interface {
	move()
}

type sayer interface {
	say()
}

type person struct {
	name string
	age  int8
}

// 使用值接收者实现接口: 类型的值和类型的指针都能够保存到接口变量中
//func (p person) move() {
//	fmt.Printf("%s在跑...\n", p.name)
//}

// 使用指针接收者实现接口：只有类型指针能够保存到接口变量中
func (p *person) move() {
	fmt.Printf("%s在跑...\n", p.name)
}

func (p *person) say() {
	fmt.Printf("%s在叫...\n", p.name)
}

func main() {
	var m mover // 定义一个mover类型的变量
	//p1 := person{ // p1是person类型的值
	//	name: "小王子",
	//	age:  18,
	//}
	p2 := &person{ // p2是person类型的指针
		name: "娜扎",
		age:  18,
	}
	//m = p1 // ? 无法赋值，因为p1是person类型的值没有实现mover接口
	m = p2
	m.move()
	fmt.Println(m)

	var s sayer // 定义一个sayer类型的变量
	s = p2
	s.say()
	fmt.Println(s)

	var a animal // 定义一个animal类型的变量
	a = p2
	a.move()
	a.say()
	fmt.Println(a)
}
