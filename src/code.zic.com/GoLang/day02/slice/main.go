package main

import "fmt"

func main() {
	// var a []int

	// b := a[:]

	// b = append(b, 100, 200)

	// fmt.Println(b) //[100 200]

	// c := []string{"北京", "上海", "广州", "深圳", "苏州"}
	// fmt.Println(c) //[北京 上海 广州 深圳 苏州]
	// c = append(c[:1], c[1:2]...)
	// fmt.Println(c) //[北京 上海]

	// j := [...]int{1, 2, 3, 4, 5, 6, 7}
	// // 基于数组得到一个切片
	// k := j[:]

	// fmt.Println(k)
	// k[0] = 100

	// fmt.Println(j, k)

	// // 区别 一个打印值的内存地址 一个打印本身的内存地址 由于切片就是切片值 所以相等
	// fmt.Printf("j:%p\n", &j)
	// fmt.Printf("k:%p\n", k)

	// 基于数组切片
	a := [5]int{1, 2, 3, 4, 5} //[1 2 3 4 5]
	b := a[:4]                 //[1 2 3 4]
	fmt.Println(a)
	fmt.Println(b)
	fmt.Printf("%T\n", b) //int

	// 切片在切片
	c := b[0:len(b)] //[1 2 3 4]
	fmt.Println(c)
	fmt.Printf("%T\n", c) //int

	// make 函数构造切片
	d := make([]int, 5, 10) //[0 0 0 0 0]
	fmt.Print(d)
	fmt.Printf("%T\n", d) //int
	// 通过len()函数获取切片长度
	fmt.Println(len(d))
	// 通过cap()函数获取切片容量
	fmt.Println(cap(d))

	// 切片的遍历
	// 根据索引遍历
	for i := 0; i < len(a); i++ {
		fmt.Println(i, a[i])
	}

	fmt.Println()
	/*
		0 1
		1 2
		2 3
		3 4
		4 5
	*/

	// for range遍历
	for k, v := range a {
		fmt.Println(k, v)
	}
	/*
		0 1
		1 2
		2 3
		3 4
		4 5
	*/

	// 切片扩容
	// 切片要初始化之后才可以使用
	var z []int //此时 它并没有申请内存

	// z = append(z, 10)

	for i := 0; i < 10; i++ {
		z = append(z, i)
		fmt.Println(z)
	}

}
