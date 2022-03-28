package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	startTime := now.UnixNano()
	sum := 0
	for i := 0; i <= 100000000; i++ {
		sum += i
	}

	endTime := time.Now().UnixNano()
	runTime := endTime - startTime
	fmt.Println(runTime) // 25488300
}
