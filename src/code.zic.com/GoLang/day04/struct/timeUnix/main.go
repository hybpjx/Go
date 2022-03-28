package main

import (
	"fmt"
	"time"
)

func timestampDemo() {
	now := time.Now()                                 //获取当前时间
	timestamp1 := now.Unix()                          //时间戳
	timestamp2 := now.UnixNano()                      //纳秒时间戳
	fmt.Printf("current timestamp1:%v\n", timestamp1) //current timestamp1:1641795975
	fmt.Printf("current timestamp2:%v\n", timestamp2) //current timestamp2:1641795975029091100
}
func timestampDemo2(timestamp int64) {
	timeObj := time.Unix(timestamp, 0)                                                  //将时间戳转为时间格式
	fmt.Println(timeObj)                                                                //2022-01-10 14:21:32 +0800 CST
	year := timeObj.Year()                                                              //年
	month := timeObj.Month()                                                            //月
	day := timeObj.Day()                                                                //日
	hour := timeObj.Hour()                                                              //小时
	minute := timeObj.Minute()                                                          //分钟
	second := timeObj.Second()                                                          //秒
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second) // 2022-01-10 14:21:32
}

func main() {
	timestampDemo()
	timestampDemo2(1641795692)
}

const (
	Nanosecond  time.Duration = 1
	Microsecond               = 1000 * Nanosecond
	Millisecond               = 1000 * Microsecond
	Second                    = 1000 * Millisecond
	Minute                    = 60 * Second
	Hour                      = 60 * Minute
)
