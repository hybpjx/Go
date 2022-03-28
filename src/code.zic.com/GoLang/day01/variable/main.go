package main

import (
	"fmt"
	"math"
)

func main() {
	var a int = 10
	var b int = 077
	var c int = 0xff

	fmt.Println(a, b, c) // 10 63 255
	// 二进制进制 以b开头
	fmt.Printf("%b\n", a) // 1010
	fmt.Printf("%b\n", b) // 111111
	fmt.Printf("%b\n", c) // 11111111
	// 八进制 以0 开头
	fmt.Printf("%o\n", a) // 12
	fmt.Printf("%o\n", b) // 77
	fmt.Printf("%o\n", c) // 377
	// 十六进制 以0x 开头
	fmt.Printf("%x\n", a) // a
	fmt.Printf("%X\n", a) // A
	fmt.Printf("%x\n", b) // 3f
	fmt.Printf("%X\n", b) // 3F
	fmt.Printf("%x\n", c) // ff
	fmt.Printf("%X\n", c) // FF

	// 求 c变量的内存地址
	fmt.Printf("%p\n", &c) // 0xc000012098

	// 浮点数的常量
	fmt.Println(math.MaxFloat32) // 3.4028234663852886e+38
	fmt.Println(math.MaxFloat64) // 1.7976931348623157e+308
}
