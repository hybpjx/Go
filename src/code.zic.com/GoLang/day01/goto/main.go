package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		for j := 0; j < 3; j++ {
			if i == 2 && j == 2 {
				goto finish
			}
			fmt.Printf("%d-----%d\n", i, j)
		}

	}

finish:
	fmt.Println("两层循环结束")
}

/*
0-----0
0-----1
0-----2
1-----0
1-----1
1-----2
2-----0
2-----1
两层循环结束
*/
