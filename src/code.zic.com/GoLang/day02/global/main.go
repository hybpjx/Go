package main

import "fmt"

//定义全局变量num
var num int64 = 10

func testNum() {
	num := 100
	fmt.Printf("num=%d\n", num) // 函数中优先使用局部变量
}

func testLocalVar2(x, y int) {
	fmt.Println(x, y) //函数的参数也是只在本函数中生效
	if x > 0 {
		z := 100 //变量z只在if语句块生效
		fmt.Println(z)
	}
	//fmt.Println(z)//此处无法使用变量z
}

type calculation func(int, int) int

func add(x, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}

func testLocalVar3() {
	for i := 0; i < 10; i++ {
		fmt.Println(i) //变量i只在当前for语句块中生效
	}
	//fmt.Println(i) //此处无法使用变量i
}

// func main() {
// 	testNum() // num=100
// 	testLocalVar2(1, 2)

// }

// func main() {
// 	var c calculation               // 声明一个calculation类型的变量c
// 	c = add                         // 把add赋值给c
// 	fmt.Printf("type of c:%T\n", c) // type of c:main.calculation
// 	fmt.Println(c(1, 2))            // 像调用add一样调用c

// 	f := add                        // 将函数add赋值给变量f1
// 	fmt.Printf("type of f:%T\n", f) // type of f:func(int, int) int
// 	fmt.Println(f(10, 20))          // 像调用add一样调用f
// }

func adder() func(int) int {
	var x int
	return func(y int) int {
		x += y
		return x
	}
}

// func main() {
// 	var f = adder()
// 	fmt.Println(f(10)) //10
// 	fmt.Println(f(20)) //30
// 	fmt.Println(f(30)) //60

// 	f1 := adder()
// 	fmt.Println(f1(40)) //40
// 	fmt.Println(f1(50)) //90
// }

// func adder2(x int) func(int) int {
// 	return func(y int) int {
// 		x += y
// 		return x
// 	}
// }
// func main() {
// 	var f = adder2(10)
// 	fmt.Println(f(10)) //20
// 	fmt.Println(f(20)) //40
// 	fmt.Println(f(30)) //70

// 	f1 := adder2(20)
// 	fmt.Println(f1(40)) //60
// 	fmt.Println(f1(50)) //110
// }

func calc(base int) (func(int) int, func(int) int) {
	add := func(i int) int {
		base += i
		return base
	}

	sub := func(i int) int {
		base -= i
		return base
	}
	return add, sub
}

func main() {
	f1, f2 := calc(10)
	fmt.Println(f1(1), f2(2)) //11 9
	fmt.Println(f1(3), f2(4)) //12 8
	fmt.Println(f1(5), f2(6)) //13 7
}
