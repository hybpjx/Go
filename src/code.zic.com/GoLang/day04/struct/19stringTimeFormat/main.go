package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now)
	// 加载时区
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println(err) // 2022-01-10 14:57:38.5306413 +0800 CST m=+0.001545501
		return
	}
	// 按照指定时区和指定格式解析字符串时间
	timeObj, err := time.ParseInLocation("2006/01/02 15:04:05", "2022/01/10 14:15:20", loc)
	if err != nil {
		fmt.Println(err) // 2022-01-10 14:15:20 +0800 CST
		return
	}
	fmt.Println(timeObj)
	fmt.Println(timeObj.Sub(now)) // -42m18.5306413s
}
