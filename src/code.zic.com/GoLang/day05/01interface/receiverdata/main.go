package main

import "fmt"

// type Mover interface {
// 	move()
// }

// type dog struct{}

// func (d dog) move() {
// 	fmt.Println("狗会移动")
// }

// func main() {
// 	var x Mover
// 	var wangcai = dog{} // 旺财是dog类型
// 	x = wangcai
// 	x.move()           // x可以接收dog类型
// 	var fugui = &dog{} // 富贵是*dog类型
// 	x = fugui          // x可以接收*dog类型
// 	x.move()
// }

type People interface {
	Speak(string) string
}

type Student struct{}

func (stu *Student) Speak(think string) (talk string) {
	if think == "sb" {
		talk = "你是个大帅比"
	} else {
		talk = "您好"
	}
	return
}

// type dog struct {
// 	name string
// }

// // 实现Sayer接口
// func (d dog) say() {
// 	fmt.Printf("%s会叫汪汪汪\n", d.name)
// }

// // 实现Mover接口
// func (d dog) move() {
// 	fmt.Printf("%s会动\n", d.name)
// }

// // Mover 接口
// type Mover interface {
// 	move()
// }

func main() {
	var peo People = &Student{}
	think := "bitch"
	fmt.Println(peo.Speak(think))

}
