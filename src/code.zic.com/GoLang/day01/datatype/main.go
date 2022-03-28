package main

import (
	"fmt"
	"strings"
)

func main() {
	s1 := "veklor"
	fmt.Println(len(s1)) //6
	// 字符串拼接
	s2 := "python"
	fmt.Println(s1 + s2)                   // veklorpython
	s3 := fmt.Sprintf("%s-----%s", s1, s2) // veklor-----python
	fmt.Println(s3)
	// 切割字符串
	ret := strings.Split(s1, "k") // [ve lor]
	fmt.Println(ret)
	// 判断是否包含
	ret2 := strings.Contains(s1, "l")
	fmt.Println(ret2) // true 返回一个bool值
	// 判断前缀 后缀
	ret3 := strings.HasPrefix(s3, "veklor")
	ret4 := strings.HasSuffix(s3, "veklor")

	fmt.Println(ret3) // true
	fmt.Println(ret4) // false

	// 求字符串的位置
	s4 := "apple"
	fmt.Println(strings.Index(s4, "l"))     // 3
	fmt.Println(strings.LastIndex(s4, "l")) // 3

	// 用join合并
	a1 := []string{"Python", "Go", "C++", "Jave"}
	fmt.Println(strings.Join(a1, "~")) // Python~Go~C++~Jave

}
