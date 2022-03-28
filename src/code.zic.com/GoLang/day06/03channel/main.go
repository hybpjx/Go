// select 多路复用
package main

import (
	"fmt"
	"math"
	"time"
)

var ch1 = make(chan string, 100)
var ch2 = make(chan string, 100)

func f1(ch chan string) {
	for i := 0; i < math.MaxInt64; i++ {
		ch <- fmt.Sprintf("f1:%d\n", i)
		time.Sleep(time.Second)
	}
}

func f2(ch chan string) {
	for i := 0; i < math.MaxInt64; i++ {
		ch <- fmt.Sprintf("f2:%d\n", i)
		time.Sleep(time.Millisecond * 50)
	}
}

func main() {

	go f1(ch1)

	go f2(ch2)

	for {
		select {
		case ret := <-ch1:
			fmt.Println(ret)
		case ret := <-ch2:
			fmt.Println(ret)
		}
	}
}
