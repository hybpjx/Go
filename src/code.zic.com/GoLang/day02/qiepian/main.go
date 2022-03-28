package main

import (
	"fmt"
)

func main() {
	// 普通列表
	var a = [...]int{
		1, 2, 3, 4, 5, 6,
	}
	fmt.Println(a)
	/*
		// 切片
		var b = []int{}
		fmt.Println(b) //[]
		for i := 1; i < len(a); i++ {
			b = append(b, i)
		}
		fmt.Println(b) //[1 2 3 4 5]
	*/

	// cap 是计算容量  len是计算大小 也称长度
	var c = []int{}
	fmt.Printf("c:%v,len:%d,cap:%d,ptr:%p\n", c, len(c), cap(c), c)
	c = append(c, 1)
	fmt.Printf("c:%v,len:%d,cap:%d,ptr:%p\n", c, len(c), cap(c), c)
	c = append(c, 1)
	fmt.Printf("c:%v,len:%d,cap:%d,ptr:%p\n", c, len(c), cap(c), c)
	c = append(c, 1)
	fmt.Printf("c:%v,len:%d,cap:%d,ptr:%p\n", c, len(c), cap(c), c)
	c = append(c, 1)
	fmt.Printf("c:%v,len:%d,cap:%d,ptr:%p\n", c, len(c), cap(c), c)
	c = append(c, 1)
	fmt.Printf("c:%v,len:%d,cap:%d,ptr:%p\n", c, len(c), cap(c), c)
	/*
	   c:[],len:0,cap:0,ptr:0xd482f8
	   c:[1],len:1,cap:1,ptr:0xc0000120d0
	   c:[1 1],len:2,cap:2,ptr:0xc0000120e0
	   c:[1 1 1],len:3,cap:4,ptr:0xc000010220
	   c:[1 1 1 1],len:4,cap:4,ptr:0xc000010220
	   c:[1 1 1 1 1],len:5,cap:8,ptr:0xc00000e300
	*/

}
