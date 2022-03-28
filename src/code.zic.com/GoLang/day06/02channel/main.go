package main

import (
	"fmt"
	"math/rand"
	"sync"
)

// 生产消费者模型
// 使用 goroutine 和channel 实现一个简易的生产者消费者模型

// 生产者： 产生随机数 math/rand

// 消费者 计算每个随机数的每个位数的和 13124123213=？

// 一个生产者 20个消费者

type item struct {
	id  int64
	num int64
}

type result struct {
	item *item
	sum  int64
}

var itemChan chan *item
var resultChan chan *result

var wg sync.WaitGroup

// 生产者
func producer(ch chan *item) {
	var id int64
	for i := 0; i < 20000; i++ {
		id++
		number := rand.Int63()
		tmp := &item{
			id:  id,
			num: number,
		}

		ch <- tmp
	}
	close(ch)
}

// 计算 位数和
func clac(num int64) int64 {

	// 123%10 =120.....3 sum  = 0+3 123/10    12
	// 12 % 10 = 1 ..... 2
	// 1 % 10 = 0 ..... 1

	var sum int64
	for num > 0 {
		sum = sum + num%10
		num = num / 10
	}
	return sum

}

// 消费者
func consumer(ch chan *item, resultChan chan *result) {

	defer wg.Done()

	for tmp := range ch {
		sum := clac((*tmp).num)
		retObj := &result{
			item: tmp,
			sum:  sum,
		}
		resultChan <- retObj
	}

}

func startWorker(n int, ch chan *item, resultChan chan *result) {
	for i := 0; i < n; i++ {
		consumer(ch, resultChan)
	}
}

func main() {
	itemChan = make(chan *item, 100000)
	resultChan = make(chan *result, 100000)
	go producer(itemChan)
	wg.Add(20)
	// 同时开启20个消费者
	startWorker(20, itemChan, resultChan)
	wg.Wait() // 等到所有生产的result 结束 再继续 再打印
	close(resultChan)
	// 打印结果
	func() {
		for ret := range resultChan {
			fmt.Printf("id:%v num:%v sum:%v\n", ret.item.id, ret.item.num, ret.sum)
		}
	}()

}
