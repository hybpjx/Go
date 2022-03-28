package main

import (
	"fmt"
)

func main() {
	// 标准写法
	var a [5]int
	a = [5]int{1, 2, 3, 4, 5}
	fmt.Println(a)

	var b [3]string
	b = [3]string{"hello", "world"}
	fmt.Println(b)
	// 合并写法
	var c [3]string = [3]string{"?", "!"}
	fmt.Println(c)
	// 简便写法
	var d = [3]int{1, 2, 3}
	fmt.Println(d)

	// 省略数字
	var e = [...]string{"I", "love", "You"}
	fmt.Println(e)

	// [123123123]的类型是[1]string
	var f = [...]string{"123123123"}
	fmt.Printf("%s的类型是%T\n", f, f)

	// 根据索引修改值
	var g = [7]int{1, 2, 3, 4, 5, 6, 7}
	var h = [7]int{6: 1}

	fmt.Println(g[3])
	fmt.Println(h)

}
