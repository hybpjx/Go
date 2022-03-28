package main

import "fmt"

func main() {
	s1 := "hello"

	byteArry := []byte(s1)

	// fmt.Println(byteArry)

	s2 := ""
	for i := len(byteArry) - 1; i >= 0; i-- {
		// 加入字符串中 需要强转
		s2 += string(byteArry[i])
	}
	fmt.Println(s2)

	// fmt.Println(s2)
	// byteArry := []byte(s1)[104 101 108 108 111]
	for i := 0; i < len(byteArry)/2; i++ {
		byteArry[i], byteArry[len(byteArry)-1-i] = byteArry[len(byteArry)-1-i], byteArry[i]
	}

	fmt.Println(string(byteArry))

}
