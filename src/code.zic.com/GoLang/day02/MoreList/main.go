package main

import (
	"fmt"
)

func main() {

	//  多维数组
	var ArrList [3][2]int

	ArrList = [3][2]int{
		[2]int{1, 2},
		[2]int{3, 4},
		[2]int{5, 6},
	}

	fmt.Println(ArrList)

	// 声明变量的同时完成赋值

	var a = [3][2]string{
		{"a", "b"},
		{"c", "d"},
		{"e", "f"},
	}

	fmt.Println(a)

	// 多维数组 除了第一层可以用... 其他都不可以 例

	var b = [...][2]int{
		{12, 34},
		{56, 78},
	}
	fmt.Println(b)
	fmt.Printf("%T\n", b) //[2][2]int

	// 多维数组的遍历

	for i := 0; i < len(b); i++ {
		for j := 0; j < len(b[i]); j++ {
			fmt.Printf("%d---%d\n", i, j)
		}
	}
	fmt.Printf("******************************************\n")
	for _, value := range b {
		for _, value2 := range value {
			fmt.Printf("%d======%d\n", value, value2)
		}
	}

}
