/*
使用函数实现一个简单的的图书管理系统
每本书都有书名  作者 价格 是否上架
用户可以在控制台添加书籍 修改书籍 打印所有的书本信息
*/

package main

import (
	"fmt"
	"os"
)

type book struct {
	bookname string
	author   string
	price    float64
	publish  bool
}

var (
	// 定义一个book指针的切片
	allbook = make([]book, 0, 200)
)

//
func NewBook(bookname, author string, price float64, publish bool) *book {
	return &book{
		bookname: bookname,
		author:   author,
		price:    price,
		publish:  publish,
	}
}

// 管理系统 header
func showMenu() {
	fmt.Println("*********************************************")
	fmt.Println("!欢迎登陆EMBS图书管理系统!")
	fmt.Println("1.添加图书")
	fmt.Println("2.修改图书信息")
	fmt.Println("3.展示所有书籍")
	fmt.Println("4.退出")
	fmt.Println("*********************************************")
	fmt.Println()

}

// 等待用户输入
func UserInput() *book {
	// 初始化一个结构体
	var add = new(book)
	fmt.Println("请根据提示输入相关内容")
	fmt.Print("请输入书名:")
	fmt.Scanln(&add.bookname)
	fmt.Print("请输入作者:")
	fmt.Scanln(&add.author)
	fmt.Print("请输入价格:")
	fmt.Scanln(&add.price)
	fmt.Print("请选择是否上架:")
	fmt.Scanln(&add.publish)
	// 返回一个结构体
	return add
}

// 添加图书 函数
func addbook() {

	// fmt.Println(*add)
	// 创建一本新书
	add := UserInput()
	book := *NewBook(add.bookname, add.author, add.price, add.publish)

	//  判断书是否存在
	for _, v := range allbook {
		if add.bookname == v.bookname {
			fmt.Printf("《%s》书已存在,请勿重复添加\n", add.bookname)
			return
		}
	}
	// 把新书添加进书架（切片）
	allbook = append(allbook, book)
	fmt.Println("添加书籍成功")
	fmt.Println()
}

// 修改图书 函数
func editbook() {
	for _, v := range allbook {
		fmt.Printf("书名:《%s》 | 作者:%s | 价格:%.2f | 是否上架:%t\n ", v.bookname, v.author, v.price, v.publish)
		fmt.Println()

		fmt.Println("请保证书名唯一性")

		edit := UserInput()
		// 遍历书籍	根据书名找到要修改的那本书
		for _, v := range allbook {
			// 如果遍历到的这本书的书名 等于 book.bookname  就更新
			if v.bookname == edit.bookname {
				v = *edit
				fmt.Println(v.bookname, "更新成功")
				return
			}
		}

		fmt.Print("******")
		fmt.Println("没有找到此书")
		fmt.Print("******")
	}
}

// 展示图书 函数
func showbook() {
	if len(allbook) == 0 {
		fmt.Println()
		fmt.Println("暂时没有数据")
		return

	}

	for _, v := range allbook {
		fmt.Printf("书名:《%s》 | 作者:%s | 价格:%.2f | 是否上架:%t\n ", v.bookname, v.author, v.price, v.publish)
		fmt.Println()
	}
}

func main() {
	for {
		showMenu()
		var input int
		fmt.Print("请选择 一个选项 请输入数字:")
		fmt.Scanf("%d\n", &input)
		switch input {
		case 1:
			addbook()
		case 2:
			editbook()
		case 3:
			showbook()
		case 4:
			fmt.Println("退出 感谢您的访问 欢迎下次使用")
			os.Exit(0)
		}
	}
}
