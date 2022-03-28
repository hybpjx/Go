package main

import (
	"fmt"
)

// func main() {

// 	// 光声明map类型，但是没有初始化 scoreMap就是初始化nill
// 	var a map[string]int
// 	fmt.Println(a == nil)

// 	// map 初始化
// 	a = make(map[string]int, 8)
// 	fmt.Println(a == nil)

// 	// 重新声明并且初始化一个变量
// 	scoreMap := make(map[string]int, 0)
// 	// 像map添加键值对
// 	scoreMap["小明"] = 100
// 	scoreMap["小王"] = 90
// 	fmt.Println(scoreMap)
// 	fmt.Println(scoreMap["小明"])
// 	fmt.Printf("a:%#v\n", scoreMap)
// 	fmt.Printf("type:%T\n", scoreMap)

// }

// func main() {
// 	userInfo := map[string]string{
// 		"username": "中东悍匪",
// 		"password": "11111",
// 	}
//     // map 是无序的 和添加顺序无关
// 	fmt.Println(userInfo)
// 	fmt.Println(userInfo["username"])
// 	fmt.Println(userInfo["password"])
// }

// func main() {
// 	scoreMap := make(map[string]int)
// 	scoreMap["张三"] = 90
// 	scoreMap["小明"] = 100
// 	// 如果key存在ok为true,v为对应的值；不存在ok为false,v为值类型的零值
// 	v, ok := scoreMap["张三"]
// 	if ok {
// 		// true
// 		fmt.Println(ok)
// 		fmt.Println(v)
// 	} else {
// 		fmt.Println("查无此人")
// 	}
// 	value, ok := scoreMap["李四"]
// 	if ok {
// 		// 不打印
// 		fmt.Println(ok)
// 		fmt.Println(value)
// 	} else {
// 		fmt.Println("查无此人")
// 	}
// }

// func main() {
// 	// 声明并且初始化
// 	a := make(map[string]int)
// 	fmt.Println(a) //map[]

// 	scoreMap := make(map[string]int)
// 	// 添加键值对
// 	scoreMap["张三"] = 90
// 	scoreMap["小明"] = 100
// 	scoreMap["娜扎"] = 60
// 	// range遍历 所有的键值
// 	for k, v := range scoreMap {
// 		fmt.Println(k, v)
// 	}
// }

// func main() {
// 	// 声明并且初始化
// 	a := make(map[string]int)
// 	fmt.Println(a) //map[]

// 	scoreMap := make(map[string]int)
// 	// 添加键值对
// 	scoreMap["张三"] = 90
// 	scoreMap["小明"] = 100
// 	scoreMap["娜扎"] = 60
// 	// range遍历 所有的键值
// 	for k := range scoreMap {
// 		fmt.Println(k)
// 	}
// }

// func main() {
// 	scoreMap := make(map[string]int)
// 	scoreMap["张三"] = 90
// 	scoreMap["小明"] = 100
// 	scoreMap["娜扎"] = 60
// 	fmt.Println(scoreMap)
// 	delete(scoreMap, "小明") //将小明:100从map中删除
// 	for k, v := range scoreMap {
// 		fmt.Println(k, v)
// 	}
// }

// import (
// 	"fmt"
// 	"math/rand"
// 	"sort"
// )

// func main() {
// 	var scoreMap = make(map[string]int, 100)

// 	for i := 0; i < 50; i++ {
// 		key := fmt.Sprintf("stu%02d", i) //生成stu开头的字符串
// 		value := rand.Intn(100)          //生成0~99的随机整数
// 		scoreMap[key] = value
// 	}

// 	for k, v := range scoreMap {
// 		fmt.Println(k, v)
// 	}

// 	// 按照k 的大小进行排序
// 	keys := make([]string, 0, 100)
// 	for k := range scoreMap {
// 		keys = append(keys, k)
// 	}
// 	// 用 sort.String 对切片进行排序
// 	sort.Strings(keys)
// 	// 1. 先取出所有的大小放入切片中
// 	// 2. 对key进行排序
// 	// 3. key排序后再对 scoremap进行排序

// 	for _, key := range keys {
// 		fmt.Println(key, scoreMap[key])
// 	}
// }

// // 类似于 [{"小明"：100}]
// func main() {
// 	// 元素为map类型的切片
// 	var mapSlice = make([]map[string]int, 8) //支持完成了切片的初始化 只完成了切片的初始化
// 	fmt.Println(mapSlice)                    //[map[] map[] map[]]...
// 	fmt.Println(mapSlice[0] == nil)          //true

// 	mapSlice[0] = make(map[string]int, 8)

// 	// panic: assignment to entry in nil map 不能给一个nill类型的map赋值
// 	mapSlice[0]["小王子"] = 100

// 	fmt.Println(mapSlice)

// }

// // 类似于{"a":["aklsjdla","asdasd","asdsad"]}
// func main() {
// 	var sliceMap = make(map[string][]int, 3) //只完成了map的初始化
// 	v, ok := sliceMap["中国"]
// 	if ok {
// 		fmt.Println(v)
// 	} else {
// 		fmt.Println("集合中没有这个国家")
// 		sliceMap["中国"] = make([]int, 8) //完成了对切片的初始化
// 		sliceMap["中国"][0] = 100
// 		sliceMap["中国"][1] = 200
// 		sliceMap["中国"][2] = 300
// 		sliceMap["中国"][3] = 400

// 	}
// 	fmt.Println(sliceMap)
// }

// // 统计 how do you do 每个单词 出现的次数
// func main() {
// 	// 1. 定义一个map集合
// 	world := "how do you do"
// 	var worldCount = make(map[string]int, 10)
// 	// 2. 字符串中都有哪些单词
// 	words := strings.Split(world, " ")

// 	fmt.Println(words) //[how do you do]

// 	for _, value := range words {
// 		// fmt.Println(value)
// 		v, ok := worldCount[value]
// 		if ok {
// 			// fmt.Println(v)
// 			// 集合中有这个单词 则 value +1
// 			worldCount[value] = v + 1
// 		} else {
// 			// 如果集合中没有这个单词 则初始值赋值为1
// 			worldCount[value] = 1
// 		}

// 	}
// 	fmt.Println(worldCount)
// 	// 3. 遍历单词做统计
// }

func main() {
	type Map map[string][]int
	m := make(Map)
	s := []int{1, 2}
	s = append(s, 3)
	fmt.Println(s)
	m["q1mi"] = s
	s = append(s[:1], s[2:]...)
	fmt.Printf("%+v\n", s)
	fmt.Printf("%+v\n", m["q1mi"])
	fmt.Println(m)
}
