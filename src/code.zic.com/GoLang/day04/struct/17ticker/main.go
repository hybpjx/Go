package main

import (
	"fmt"
	"time"
)

func tickDemo() {
	ticker := time.Tick(time.Second) //定义一个1秒间隔的定时器
	for i := range ticker {
		fmt.Println(i) //每秒都会执行的任务
	}
}

func main() {
	tickDemo()
}

/*
2022-01-10 14:50:02.8798808 +0800 CST m=+1.003672301
2022-01-10 14:50:03.8817002 +0800 CST m=+2.005491701
2022-01-10 14:50:04.8790619 +0800 CST m=+3.002853401
2022-01-10 14:50:05.8789413 +0800 CST m=+4.002732801
2022-01-10 14:50:06.8793836 +0800 CST m=+5.003175101
...
*/
