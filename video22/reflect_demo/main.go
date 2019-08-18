package main

import (
	"fmt"
	"reflect"
)

// reflect demo

func reflectType(x interface{}) {
	// 我是不知道别人调用我这个函数的时候会传进来什么类型的变量
	// 1. 方式1：通过类型断言
	// 2. 方法2：借助反射
	obj := reflect.TypeOf(x)
	fmt.Println(obj, obj.Name(), obj.Kind())
	fmt.Printf("%T\n", obj) // *reflect.rtype
}

func reflectValue(x interface{}) {
	v := reflect.ValueOf(x)
	fmt.Printf("%v,%T\n", v, v)
	k := v.Kind() // 拿到值对应的类型种类
	fmt.Println(k)
	// 如何得到一个传入时候类型的变量
	switch k {
	case reflect.Float32:
		// 把反射取到的值转换成一个float32类型的变量
		ret := float32(v.Float())
		fmt.Printf("%v %T\n", ret, ret)
	case reflect.Int32:
		ret := int32(v.Int())
		fmt.Printf("%v %T\n", ret, ret)
	}
}

func reflectSetValue(x interface{}) {
	v := reflect.ValueOf(x)
	// Elem()用来根据指针取对应的值
	k := v.Elem().Kind()
	switch k {
	case reflect.Int32:
		v.Elem().SetInt(100)
	case reflect.Float32:
		v.Elem().SetFloat(3.21)
	}
}

type Cat struct{}

type Dog struct{}

func main() {
	//var a float32 = 1.234
	//reflectType(a)
	//var b int8 = 10
	//reflectType(b)
	//// 结构体类型
	//var c Cat
	//reflectType(c)
	//var d Dog
	//reflectType(d)
	//// slice
	//var e []int
	//reflectType(e)
	//var f []string
	//reflectType(f)

	// valueOf
	//var aa int32 = 100
	//reflectValue(aa)
	//var bb float32 = 1.234
	//reflectValue(bb)

	// set value
	var aaa int32 = 10
	reflectSetValue(&aaa)
	fmt.Println(aaa)
}
