package main

import "fmt"

type Sayer interface {
	say()
}

type cat struct {
}

func (c cat) say() {
	fmt.Println("猫在喵喵")
}

type dog struct {
}

func (d dog) say() {
	fmt.Println("狗在汪汪")
}

func main() {
	var x Sayer // 声明一个Sayer类型的变量x
	a := cat{}  // 实例化一个cat
	b := dog{}  // 实例化一个dog
	x = a       // 可以把cat实例直接赋值给x
	x.say()     // 喵喵喵
	x = b       // 可以把dog实例直接赋值给x
	x.say()     // 汪汪汪
}
