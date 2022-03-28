package main

import "fmt"

func SwithchDemo() {
	finger := 3
	switch finger {
	case 1:
		fmt.Println("大拇指")
	case 2:
		fmt.Println("食指")
	case 3:
		fmt.Println("中指")
	case 4:
		fmt.Println("无名指")
	case 5:
		fmt.Println("小拇指")
	default:
		fmt.Println("脚趾")
	}
}

func JSwitchdemo() {

	switch n := 7; n {
	case 1, 3, 5, 7, 9:
		fmt.Println("奇数")
	case 2, 4, 6, 8:
		fmt.Println("偶数")
	default:
		fmt.Printf("是零 或者是%d", &n)
	}
}

func Switchdemo1() {
	s := "a"
	switch {
	case s == "a":
		fmt.Println("a")
		fallthrough
	case s == "b":
		fmt.Println("b")
	case s == "c":
		fmt.Println("c")
	default:
		fmt.Println("alskjdklsajldsja")
	}
}

func main() {

	if age := 17; age > 18 {
		fmt.Println("澳门皇家线上赌场")
	} else if age < 18 {
		fmt.Println("未成年禁止入内")
	} else {
		fmt.Println("好，你是煞笔")
	}
	SwithchDemo()

	JSwitchdemo()

	Switchdemo1()
}
