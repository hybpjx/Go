package main

import "fmt"

type Cat struct{}

func (c Cat) Say() string {
	return "喵喵喵~"
}

type Dog struct{}

func (d Dog) Say() string {
	return "汪汪汪~"
}

type Sheep struct{}

func (s Sheep) Say() string {
	return "咩咩咩~"
}

// func MakeCatHungry(c Cat) {
// 	fmt.Println("🐱饿了：", c.Say())
// }

// func MakeSheepHungry(s Sheep) {
// 	fmt.Println("🐑饿了：", s.Say())
// }

type sayer interface {
	Say() string
}

func MakeHungry(s sayer) {
	fmt.Println("饿了：", s.Say())
}

func main() {
	// c := Cat{}
	// fmt.Println("猫:", c.Say())
	// d := Dog{}
	// fmt.Println("狗:", d.Say())
	// s := Sheep{}
	// fmt.Println("🐏:", s.Say())

	// MakeCatHungry(c)

	// MakeSheepHungry(s)

	var c Cat

	MakeHungry(c)

	var s Sheep
	MakeHungry(s)
}
