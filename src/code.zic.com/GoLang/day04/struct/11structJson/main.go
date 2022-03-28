package main

import (
	"encoding/json"
	"fmt"
)

//Student 学生
type Student struct {
	ID     int
	Gender string
	Name   string
}

func main() {
	stu1 := new(Student)
	stu1.ID = 1
	stu1.Gender = "男"
	stu1.Name = "豪杰"

	// 序列化 把 Go语言里面的数据转换成字符串  遵循json格式的字符串
	v, err := json.Marshal(stu1)
	fmt.Println(err) // <nil>
	if err != nil {
		fmt.Println(err)
		fmt.Println("格式化出错了")
	}
	// 切片类型的数据 [] bytes
	fmt.Println(v) // [123 34 73 68 34 58 49 44 34 71 101 110 100 101 114 34 58 34 231 148 183 34 44 34 78 97 109 101 34 58 34 232 177 170 230 157 176 34 125]
	// 强转换成字符串
	fmt.Println(string(v))         //{"ID":1,"Gender":"男","Name":"豪杰"}
	fmt.Printf("%#v\n", string(v)) // "{\"ID\":1,\"Gender\":\"男\",\"Name\":\"豪杰\"}"

	// 反序列化 把json格式的字符串转化成Go语言里的字符串

	str := string(v)
	var stu2 = &Student{}
	json.Unmarshal([]byte(str), stu2)
	fmt.Println(stu2)        //&{1 男 豪杰}
	fmt.Printf("%T\n", stu2) // *main.Student
}
