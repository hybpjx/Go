package main

import "fmt"

func justifyType(x interface{}) {
	switch v := x.(type) {
	case string:
		fmt.Printf("x is a string，value is %v\n", v)
	case int:
		fmt.Printf("x is a int is %v\n", v)
	case bool:
		fmt.Printf("x is a bool is %v\n", v)
	default:
		fmt.Println("unsupport type！")
	}
}

func main() {
	var x interface{}
	x = "Hello 沙河"
	v, ok := x.(string)
	fmt.Println(v, ok)
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("类型断言失败")
	}

}
