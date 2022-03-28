package main

import "fmt"

// const (
// 	_  = iota
// 	KB = 1 << (10 * iota)
// 	MB = 1 << (10 * iota)
// 	GB = 1 << (10 * iota)
// 	TB = 1 << (10 * iota)
// 	PB = 1 << (10 * iota)
// )

// const (
// 	a, b = iota + 1, iota + 2
// 	c, d
// 	e, f
// 	g, h
// 	ab   = 1
// 	z, m = iota + 1, iota + 2
// )

const pi = 3.1415926

func main() {
	// 常量不允许修改
	var pi = 3.1415926123213213
	fmt.Println(pi)

	// fmt.Println(KB, MB, GB, TB, PB)
	// fmt.Println(a, b, c, d, e, f, g, h, ab, z, m)
}

// iota
// 0. const 声明如果不写， 默认和上一行一样
// 1. 遇到const iota 就初始化为零
// 2. const 没新增一行变量生命 iota 就递增为1
