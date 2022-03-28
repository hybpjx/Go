package main

import "fmt"

func main() {
	s1 := "hello"
	for i := 0; i < len(s1); i++ {
		fmt.Printf("%c\n", s1[i])

	}

	for k, v := range s1 {
		fmt.Printf("%d,%c\n", k, v)
	}

	// 强制类型转换

	s2 := "big"
	byteArray := []byte(s2)
	fmt.Println(byteArray)

	byteArray[0] = 'p'
	// 将字节数组强制转换为字符串类型
	s2 = string(byteArray)

	fmt.Println(s2)

}
