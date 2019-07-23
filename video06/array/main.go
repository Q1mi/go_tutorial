package main

import "fmt"

// 数组相关内容
func main() {
	// var a [3]int
	// var b [4]int
	// // a = b
	// fmt.Println(a)
	// fmt.Println(b)
	// 数组的初始化
	// 1. 定义时使用初始值列表的方式初始化
	// var cityArray = [4]string{"北京", "上海", "广州", "深圳"}
	// fmt.Println(cityArray)
	// fmt.Println(cityArray[0])
	// fmt.Println(cityArray[3])
	// // 2. 编译器推导数组的长度
	// var boolArray = [...]bool{true, false, false, true, false}
	// fmt.Println(boolArray)

	// // 3. 使用索引值方式初始化
	// var langArray = [...]string{1: "Golang", 3: "Python", 7: "Java"}
	// fmt.Println(langArray)
	// fmt.Printf("%T\n", langArray)

	// 数组的遍历
	// var cityArray = [4]string{"北京", "上海", "广州", "深圳"}
	// 1. for循环遍历
	// for i := 0; i < len(cityArray); i++ {
	// 	fmt.Println(cityArray[i])
	// }
	// 2. for range遍历
	// for _, value := range cityArray {
	// 	fmt.Println(value)
	// }

	// 二维数组
	// cityArray := [...][2]string{
	// 	{"北京", "西安"},
	// 	{"上海", "杭州"},
	// 	{"重庆", "成都"},
	// 	{"广州", "深圳"},
	// }
	// fmt.Println(cityArray)
	// fmt.Println(cityArray[2][0])
	// // 二维数组的遍历
	// for _, v1 := range cityArray {
	// 	fmt.Println(v1)
	// 	for _, v2 := range v1 {
	// 		fmt.Println(v2)
	// 	}
	// }

	// 数组是值类型
	x := [3][2]int{
		{1, 2},
		{3, 4},
		{5, 6},
	}
	fmt.Println(x)
	f1(x)
	fmt.Println(x)
	y := x
	y[0][0] = 1000
	fmt.Println(x)
}

func f1(a [3][2]int) {
	a[0][0] = 100
}
