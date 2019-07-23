package main

import "fmt"

// 结构体的继承

type Animal struct {
	name string
}

func (a *Animal) move() {
	fmt.Printf("%s会动~\n", a.name)
}

type Dog struct {
	Feet    int8
	*Animal // 匿名嵌套，而且嵌套的是一个结构体指针
}

func (d *Dog) wang() {
	fmt.Printf("%s会汪汪汪~\n", d.name)
}

func main() {
	d1 := &Dog{
		Feet: 4,
		Animal: &Animal{
			name: "乐乐",
		},
	}
	d1.wang() // 狗会wang
	d1.move() // 狗会动
}
