package main

import "fmt"

// for循环
func main() {
	// 1. for循环的基本写法
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(i)
	// }
	// 2. 省略初识语句，但是必须保留初始语句后面的分号
	// var i = 0
	// for ; i < 10; i++ {
	// 	fmt.Println(i)
	// }
	// 3. 省略初始语句和结束语句
	// var i = 10
	// for i > 0{
	// 	fmt.Println(i)
	// 	i--
	// }
	// 4. 无限循环
	// for {
	// 	fmt.Println("Hello 沙河！")
	// }
	// 5. break 跳出for循环
	// for i := 0; i < 5; i++ {
	// 	fmt.Println(i)
	// 	if i == 3 {
	// 		break
	// 	}
	// }
	// 6. continue 继续下一次循环
	for i := 0; i < 5; i++ {
		if i == 3 {
			continue//跳过本次for循环，继续下一次循环
		}
		fmt.Println(i)
	}
}
