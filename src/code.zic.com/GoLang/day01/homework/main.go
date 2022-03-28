package main

import "fmt"

// func JiuJiu() {

// 	for i := 1; i <= 9; i++ {
// 		for j := 1; j <= i; j++ {
// 			// fmt.Println(i,j)
// 			fmt.Printf("%d*%d=%d\t", j, i, i*j)

// 		}
// 		fmt.Println()
// 	}
// }

func Su() {

	for i := 200; i < 1000; i++ {
		for j := 2; j < i; j++ {
			if i%j == 0 {
				fmt.Println(j)
			}

		}

	}
}

// const pi = 3.1415926

func main() {
	// JiuJiu()
	Su()
}
