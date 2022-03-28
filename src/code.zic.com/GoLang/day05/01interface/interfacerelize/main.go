package main

import "fmt"

// Washingmachine 接口
type Washingmachiner interface {
	wash() // 洗
	dry()  // 甩干
}

type Haier struct {
	name  string
	price float64
	mode  string
}

func (h Haier) wash() {
	fmt.Println("海尔洗衣机洗衣服")
}

func (h Haier) dry() {
	fmt.Println("海尔洗衣机甩干")
}

func main() {
	var wm Washingmachiner

	h1 := Haier{
		name:  "海尔",
		price: 998.8,
		mode:  "滚筒式",
	}
	fmt.Printf("h1的类型：%T \n wm的类型：%T \n", h1, wm)

	wm = h1         //接口是一种类型 一种抽象的类型
	fmt.Println(wm) // {海尔 998.8 滚筒式}
	wm.wash()       // 海尔洗衣机洗衣服
	wm.dry()        // 海尔洗衣机甩干

}
