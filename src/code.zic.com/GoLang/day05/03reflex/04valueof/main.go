package main

import (
	"fmt"
	"reflect"
)

func reflectValue(x interface{}) {
	v := reflect.ValueOf(x)
	k := v.Kind()
	switch k {
	case reflect.Int64:
		// v.Int()从反射中获取整型的原始值，然后通过int64()强制类型转换
		fmt.Printf("type is int64, value is %d\n", int64(v.Int()))
	case reflect.Float32:
		// v.Float()从反射中获取浮点型的原始值，然后通过float32()强制类型转换
		fmt.Printf("type is float32, value is %f\n", float32(v.Float()))
	case reflect.Float64:
		// v.Float()从反射中获取浮点型的原始值，然后通过float64()强制类型转换
		fmt.Printf("type is float64, value is %f\n", float64(v.Float()))
	}
}
func main() {
	var a float32 = 3.14
	var b float64 = 3.14
	var c int64 = 100
	reflectValue(a) // type is float32, value is 3.140000
	reflectValue(b) // type is float64, value is 3.140000
	reflectValue(c) // type is int64, value is 100
	// 将int类型的原始值转换为reflect.Value类型
	d := reflect.ValueOf(10)
	fmt.Printf("type d :%T\n", d) // type d :reflect.Value
	fmt.Printf("value :%v\n", d)  // value :10
}
