package main

import "fmt"

// Mover 接口
type Mover interface {
	move()
}

type dog struct {
	name string
}

type car struct {
	brand string
}

// dog类型实现Mover接口
func (d dog) move() {
	fmt.Printf("%s会跑\n", d.name)
}

// car类型实现Mover接口
func (c car) move() {
	fmt.Printf("%s速度70迈\n", c.brand)
}

func main() {
	var m Mover

	var d = dog{"旺财"}

	m = d
	m.move()
	var c = car{"奔驰"}
	m = c
	m.move()

}
