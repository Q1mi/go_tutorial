package main

import "fmt"

// 嵌套结构体

type Address struct {
	Province string
	City     string
}

type Person struct {
	Name    string
	Gender  string
	Age     int
	Address // 嵌套另外一个结构体
}

func main() {
	p1 := Person{
		Name:   "小王子",
		Gender: "男",
		Age:    18,
		Address: Address{
			Province: "山东",
			City:     "威海",
		},
	}
	fmt.Printf("%#v\n", p1)
	fmt.Println(p1.Name, p1.Gender, p1.Age, p1.Address)
	fmt.Println(p1.Address.Province) // 通过嵌套的匿名结构体字段访问其内部的字段
	fmt.Println(p1.Province)         // 直接访问匿名结构体中的字段
}
