package main

import (
	"fmt"
	"sort"
)

func main() {
	// 这是一个数组
	var arra = [...]int{
		3, 7, 9, 8, 1,
	}
	// 由于接受的参数必须是一个切片 所以转换一下
	sort.Ints(arra[:])

	fmt.Println(arra)
}
