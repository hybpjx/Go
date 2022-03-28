package main

import "fmt"

type person struct {
	name string
	age  int
	city string
}

func main() {

	var p2 = new(person)
	p2.name = "小王子"
	p2.age = 28
	p2.city = "上海"
	fmt.Printf("p2=%#v\n", p2) //p2=&main.person{name:"小王子", city:"上海", age:28}

	p3 := &person{}
	fmt.Printf("%T\n", p3)     //*main.person
	fmt.Printf("p3=%#v\n", p3) //p3=&main.person{name:"", age:0, city:""}
	p3.name = "七米"
	p3.age = 30
	p3.city = "成都"
	fmt.Printf("p3=%#v\n", p3) //p3=&main.person{name:"七米", age:30, city:"成都"}

}
