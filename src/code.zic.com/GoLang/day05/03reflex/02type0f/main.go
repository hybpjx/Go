package main

import (
	"fmt"
	"reflect"
)

func reflectType(x interface{}) {
	v := reflect.TypeOf(x)
	fmt.Printf("type: %v\n", v)
}

func main() {
	var a int = 10
	reflectType(a)
	var b string = "hello，world"
	reflectType(b)
}
