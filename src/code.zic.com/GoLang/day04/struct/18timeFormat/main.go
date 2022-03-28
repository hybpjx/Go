package main

import (
	"fmt"
	"time"
)

func formatDemo() {
	now := time.Now()
	// 格式化的模板为Go的出生时间2006年1月2号15点04分 Mon Jan
	// 24小时制
	fmt.Println(now.Format("2006-01-02 15:04:05.000 Mon Jan")) // 2022-01-10 14:54:43.744 Mon Jan
	// 12小时制
	fmt.Println(now.Format("2006-01-02 03:04:05.000 PM Mon Jan")) //2022-01-10 02:54:43.744 PM Mon Jan
	fmt.Println(now.Format("2006/01/02 15:04"))                   //2022/01/10 14:54
	fmt.Println(now.Format("15:04 2006/01/02"))                   //14:54 2022/01/10
	fmt.Println(now.Format("2006/01/02"))                         // 2022/01/10

	now1 := time.Now()
	later := now1.Add(time.Hour)                          // 当前时间加1小时后的时间
	fmt.Println(later.Format("2006/01/02——————15:04:05")) // 2022/01/10——————15:54:43

}

func main() {
	formatDemo()
}
