package main

import "fmt"

// 方法的定义实例

// Person 是一个结构体
type Person struct {
	name string
	age  int8
}

// NewPerson 是一个Person类型的构造函数
func NewPerson(name string, age int8) *Person {
	return &Person{
		name: name,
		age:  age,
	}
}

// Dream 是为Person类型定义方法
func (p Person) Dream() {
	fmt.Printf("%s的梦想是学好Go语言！\n", p.name)
}

// SetAge 是一个修改年龄的方法
// 指针接收者指的就是接收者的类型是指针
func (p *Person) SetAge(newAge int8) {
	p.age = newAge
}

// SetAge2 是一个使用值接收者来修改年龄的方法
func (p Person) SetAge2(newAge int8) {
	p.age = newAge
}

func main() {
	p1 := NewPerson("沙河娜扎", int8(18))
	// (*p1).Dream()
	// p1.Dream()

	// 调用修改年龄的方法
	fmt.Println(p1.age)
	p1.SetAge(int8(20))
	fmt.Println(p1.age)
	p1.SetAge2(int8(30))
	fmt.Println(p1.age) // 20?
}
