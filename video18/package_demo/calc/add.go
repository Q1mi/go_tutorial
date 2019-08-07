package calc

import (
	"code/video18/package_demo/snow"
	"fmt"
)

// 标识符首字母大写表示对外可见。
// 通常不对外的标识符都是要首字母小写。

// Name定 义一个测试的全局变量
var Name = "沙河小王子"

// Add 是一个计算两个int类型数据和的函数
func Add(x, y int) int {
	snow.Snow()
	Sub(x, y) //同一个包中的不同文件可以直接调用
	return x + y
}

// init函数在包导入的时候自动执行
// init函数没有参数也没有返回值
func init() {
	fmt.Println("calc.init()")
	fmt.Println(Name)
}
