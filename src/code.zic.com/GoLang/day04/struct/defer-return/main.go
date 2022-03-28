package main

import "fmt"

// func f1() int {
// 	// 声明 i
// 	var i int
// 	defer func() {
// 		// i+1
// 		i++
// 		fmt.Println("f11: ", i)
// 	}()

// 	i = 1000
// 	return i
// }

// func f2() (i int) {
// 	defer func() {
// 		i++
// 		fmt.Println("f21: ", i)
// 	}()
// 	i = 1000
// 	return i
// }
// func main() {
// 	fmt.Println("f1 result: ", f1())
// 	fmt.Println("f2 result: ", f2())
// }

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}
func main() {
	x := 1
	y := 2
	defer calc("AA", x, calc("A", x, y))
	x = 10
	defer calc("BB", x, calc("B", x, y))
	y = 20
}
