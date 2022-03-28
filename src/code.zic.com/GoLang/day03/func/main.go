package main

import (
	"fmt"
)

func f1() {
	fmt.Println("hello,world")
}

func f2(name string) {
	fmt.Println("hello", name)
}

func f3(names ...string) {
	fmt.Println(names, "你好") // []string
}

func f4() string {
	return "hello"
}

func f5(a, b int) (int, int) {
	return a + b, a - b
}

func f6(a, b int) (sum int, sub int) {
	sum = a + b
	sub = a - b
	return
}

func main() {
	// f2("小王")

	// f3("小王", "小李", "小张")

	// f4()

	fmt.Println(f5(3, 4))
	fmt.Println(f6(3, 4))
	func(name string) {
		fmt.Println("hello", name)
	}("沙河")

}
