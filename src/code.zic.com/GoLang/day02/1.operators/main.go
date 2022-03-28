package main

import "fmt"

func main() {
	// 算术运算符
	n1 := 10
	n2 := 3
	fmt.Println(n1 + n2)
	fmt.Println(n1 - n2)
	fmt.Println(n1 * n2)
	fmt.Println(n1 / n2)
	fmt.Println(n1 % n2)
	n1++
	fmt.Println(n1)
	n2--
	fmt.Println(n2)

	// 关系运算符
	n3 := 10
	n4 := 3
	fmt.Println(n3 == n4)
	fmt.Println(n3 != n4)
	fmt.Println(n3 >= n4)
	fmt.Println(n3 <= n4)

	// 逻辑运算符
	a := true
	b := false
	fmt.Println(a && b)
	fmt.Println(a || b)
	fmt.Println(!a)
	fmt.Println(!b)

	// 位运算符
	// 1101 13             11 3
	fmt.Printf("13 的二进制%b\n", 13)
	fmt.Printf("3 的二进制%b\n", 3)

	// & 两个对应的二进制位为1 才为1
	fmt.Println(13 & 3)
	// | 两个对应的二进制位有一个为1 就为1
	fmt.Println(13 | 3)
	fmt.Println(13 & 3)
	// ^  两个数对应的二进制位不一样 就为1
	fmt.Println(13 ^ 3)
	// <<
	fmt.Println(3 << 10) //110000000000
	// >>
	fmt.Println(3 >> 1)

}
