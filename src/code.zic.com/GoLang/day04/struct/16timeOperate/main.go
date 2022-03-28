package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	later := now.Add(time.Hour) // 当前时间加1小时后的时间
	fmt.Println(later)          // 2022-01-10 15:29:53.9950449 +0800 CST m=+3600.001577101
}

type Time struct {
}

func (t Time) Sub(u Time) time.Duration

func (t Time) Equal(u Time) bool
