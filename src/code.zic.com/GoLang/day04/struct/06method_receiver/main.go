package main

import "fmt"

//Person 结构体
type Person struct {
	name string
	age  int8
}

//NewPerson 构造函数
func NewPerson(name string, age int8) *Person {
	return &Person{
		name: name,
		age:  age,
	}
}

//Dream Person做梦的方法
func (p Person) Dream() {
	fmt.Printf("%s的梦想是学好Go语言！\n", p.name)
}

// func main() {
// 	p1 := NewPerson("小王子", 25)
// 	p1.Dream()
// 	// p2 := Dream() 无法直接调用
// 	// fmt.Println(p2)
// }

// // SetAge 设置p的年龄
// // 使用指针接收者
// func (p *Person) SetAge(newAge int8) {
// 	p.age = newAge
// }

// func main() {
// 	p1 := NewPerson("小王子", 25)
// 	fmt.Println(p1.age) // 25
// 	p1.SetAge(30)
// 	fmt.Println(p1.age) // 30
// }

// // SetAge2 设置p的年龄
// // 使用值接收者
// func (p Person) SetAge2(newAge int8) {
// 	p.age = newAge
// }

// func main() {
// 	p1 := NewPerson("小王子", 25)
// 	p1.Dream()
// 	fmt.Println(p1.age) // 25
// 	p1.SetAge2(30)      // (*p1).SetAge2(30)
// 	fmt.Println(p1.age) // 25
// }

//MyInt 将int定义为自定义MyInt类型
type MyInt int

//SayHello 为MyInt添加一个SayHello的方法
func (m MyInt) SayHello() {
	fmt.Println("Hello, 我是一个int。")
}
func main() {
	var m1 MyInt
	m1.SayHello() //Hello, 我是一个int。
	m1 = 100
	fmt.Printf("%#v  %T\n", m1, m1) //100  main.MyInt
}
