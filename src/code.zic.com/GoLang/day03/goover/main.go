package main

import "fmt"

func main() {
	// 数组是值类型
	// 修改数组的实例
	a1 := [3]int{1, 2, 3}
	b1 := a1 //a1和b1 此时都有自己对应的 [1,2,3,]

	b1[0] = 20
	fmt.Println(a1)
	fmt.Println(b1)

	// 切片
	// 切片是引用累心
	var c1 = []int{1, 2, 3}
	d1 := c1

	d1[0] = 20
	fmt.Println(c1)
	fmt.Println(d1)

	// 1. 声明类型时初始化

	/*
	   切片三要素：
	   1. 第一个元素在底层数组的位置
	   2. 大小（len） 指的是当前切片中元素的数量
	   3. 容量（cap） 指的是底层数组能容纳元素的个数
	*/
	var x = []int{1, 2, 3}
	fmt.Println(x, len(x), cap(x))
	x2 := x[:2]
	x2[0] = 1999
	fmt.Println(x, len(x), cap(x))
	fmt.Println(x2, len(x2), cap(x2))

	// 切片的追加
	a := []int{1, 2, 3}
	fmt.Println(len(a), cap(a))
	a = append(a[:], 18)
	fmt.Println(len(a), cap(a))

	// 切片的复制
	b := []string{"三国演义", "西游记", "水浒传", "红楼梦"}
	c := make([]string, 4)
	copy(c, b)
	b[1] = "白蛇传"
	fmt.Println(b)
	fmt.Println(c)

	// 切片的删除
	e := []int{1, 2, 3}
	e = append(e[:1], e[2:]...)
	fmt.Println(e)

}
