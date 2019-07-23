package main

import (
	"fmt"

	"os"
)

// 学员信息管理系统

// 需求：
// 1. 添加学员信息
// 2. 编辑学员信息
// 3. 展示所有学员信息

func showMenu() {
	fmt.Println("欢迎来到学员信息管理系统")
	fmt.Println("1. 添加学员")
	fmt.Println("2. 编辑学员信息")
	fmt.Println("3. 展示所有学员信息")
	fmt.Println("4. 退出系统")
}

// 获取用户输入的学员信息
func getInput() *student {
	var (
		id    int
		name  string
		class string
	)
	fmt.Println("请按要求输入学员信息")
	fmt.Print("请输入学员的学号：")
	fmt.Scanf("%d\n", &id)
	fmt.Print("请输入学员的姓名：")
	fmt.Scanf("%s\n", &name)
	fmt.Print("请输入学员的班级：")
	fmt.Scanf("%s\n", &class)
	// 就能拿到用户输入的学员的所有信息
	stu := newStudent(id, name, class) // 调用student的构造函数造一个学生
	return stu
}

func main() {

	sm := newStudentMgr()
	for {
		// 1. 打印系统菜单
		showMenu()
		// 2. 等待用户选择要执行的选项
		var input int
		fmt.Print("请输入你要操作的序号：")
		fmt.Scanf("%d\n", &input)
		fmt.Println("用户输入的是：", input)
		// 3. 执行用户选择的动作
		switch input {
		case 1:
			// 添加学员
			stu := getInput()
			sm.addStudent(stu)
		case 2:
			// 编辑学员
			stu := getInput()
			sm.modifyStudent(stu)
		case 3:
			// 展示所有学员
			sm.showStudent()
		case 4:
			// 退出
			os.Exit(0)
		}
	}
}
