package main

import "fmt"

// func main() {
// 	// var a int
// 	// b := &a
// 	// // b的值
// 	// fmt.Printf("b=%v\n", b)
// 	// // 类型
// 	// fmt.Printf("Type=%T\n", b)
// 	// c := 100
// 	// // c的内存地址
// 	// fmt.Println(&c)
// 	// b = &c
// 	// // b的内存地址
// 	// fmt.Println(b)
// 	// // b指向的哪个值
// 	// fmt.Println(*b)

// 	// a := 10
// 	// b := &a
// 	// fmt.Printf("a:%d ptr:%p\n", a, &a) // a:10 ptr:0xc00001a078
// 	// fmt.Printf("b:%p type:%T\n", b, b) // b:0xc00001a078 type:*int
// 	// fmt.Println(&b)                    // 0xc00000e018
// 	//指针取值
// 	a := 10
// 	b := &a // 取变量a的地址，将指针保存到b中
// 	fmt.Printf("type of b:%T\n", b)  // type of b:*int
// 	c := *b // 指针取值（根据指针去内存取值）
// 	fmt.Printf("type of c:%T\n", c) // type of c:int
// 	fmt.Printf("value of c:%v\n", c) // value of c:10
// }

// 定义一个修改数组第一个元素为100的函数
func modifyArray(a1 [3]int) {
	a1[0] = 100
}

// 接收的参数是一个数组的指针
func modifyArray2(a1 *[3]int) {
	a1[0] = 100
}

func main() {
	// 指针的应用
	a := [3]int{1, 2, 3}
	fmt.Println(a) //[1 2 3]
	modifyArray(a)
	fmt.Println(a) //[1 2 3]

	modifyArray2(&a)
	fmt.Println(a)

}
