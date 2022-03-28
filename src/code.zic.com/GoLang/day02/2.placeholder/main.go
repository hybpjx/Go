package main

import "fmt"

func main() {
	var a = "hello"
	var b = 1111
	fmt.Println(a)
	// 不知道什么类型 用v
	fmt.Printf("%v,world\n", a)
	// 打印类型 用T
	fmt.Printf("%T,类型\n", a)
	// 转义符
	fmt.Printf("%d %%/\n", b)
}
