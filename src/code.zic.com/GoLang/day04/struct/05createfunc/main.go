package main

import "fmt"

type person struct {
	name, city string
	age        int8
}

func newPerson(name, city string, age int8) *person {
	return &person{
		name: name,
		city: city,
		age:  age,
	}
}

func main() {

	p := newPerson("小王", "北京", 18)
	fmt.Println(p)
}
