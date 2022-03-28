package main

import (
	"fmt"
	"os"
)

// 展示页面
func showInfo() {
	fmt.Println("欢迎使用学生管理系统")
	fmt.Println("1.增加一名学生")
	fmt.Println("2.修改一名学生")
	fmt.Println("3.删除一名学生")
	fmt.Println("4.查询所有学生")
	fmt.Println("5.退出")
}

func userInput() *Student {
	var (
		name, class string
		id, age     int
	)
	fmt.Println("请按提示输入录入信息")
	fmt.Print("请输入id:")
	fmt.Scanf("%d\n", &id)
	fmt.Print("请输入姓名:")
	fmt.Scanf("%s\n", &name)
	fmt.Print("请输入年龄:")
	fmt.Scanf("%d\n", &age)
	fmt.Print("请输入班级:")
	fmt.Scanf("%s\n", &class)
	newStu := NewStudent(id, age, name, class)
	return newStu
}

func main() {
	stuMgr := *NewStudentMgr()

	for {
		showInfo()
		// 定义一个输入的变量
		var input int
		// 调用函数
		fmt.Print("请输入你要的操作:")
		// 一定要用指针 不然传入的是0
		fmt.Scanf("%d\n", &input)
		fmt.Println(input)
		switch input {
		case 1:
			newstu := userInput()
			stuMgr.AddStudent(newstu)
		case 2:
			newstu := userInput()
			stuMgr.EditStudent(newstu)
		case 3:
			newstu := userInput()
			stuMgr.DelStudent(newstu)
		case 4:
			stuMgr.ShowStudent()
		case 5:
			fmt.Println("即将退出.......")
			os.Exit(0)
		default:
			fmt.Println("输入无效")
		}
	}

}
