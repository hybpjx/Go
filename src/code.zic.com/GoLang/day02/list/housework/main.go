package main

import "fmt"

func main() {
	// 求列表长度
	var ArrList = [...]int{1, 3, 5, 7, 8}
	fmt.Println(ArrList)
	var sum int

	for _, value := range ArrList {
		sum += value
	}
	fmt.Println(sum)

	// // 求索引下标等于8 的所有索引

	// for index, value := range ArrList {
	// 	// fmt.Println(index, value)
	// 	other := 8 - value

	// 	for index2, value := range ArrList {
	// 		if value == other {
	// 			fmt.Println(index, index2)
	// 		}
	// 	}

	// }

	for index, value := range ArrList {
		// fmt.Println(index, value)
		other := 8 - value
		for index2 := index + 1; index2 < len(ArrList); index2++ {
			if ArrList[index2] == other {
				fmt.Println(index, index2)
			}
		}
	}

}
