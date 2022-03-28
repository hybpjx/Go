package main

import "fmt"

type student struct {
	name string
	age  int
}

func main() {
	m := make(map[string]*student)
	// fmt.Printf("%T\n", m) //map[string]*main.student
	stus := []student{
		{name: "小王子", age: 18},
		{name: "娜扎", age: 23},
		{name: "大王八", age: 9000},
	}

	// print(stus, "\n") //[3/3]0xc00007fec8
	for _, stu := range stus {
		m[stu.name] = &stu
		// fmt.Println(stu)
		/*
			{小王子 18}
			{娜扎 23}
			{大王八 9000}
		*/
	}
	// fmt.Println(m) //map[大王八:0xc000096060 娜扎:0xc000096060 小王子:0xc000096060]
	for k, v := range m {
		// fmt.Println(k, v)
		/*
			小王子 &{大王八 9000}
			娜扎 &{大王八 9000}
			大王八 &{大王八 9000}
		*/
		fmt.Println(k, "=>", v.name)
	}
}
