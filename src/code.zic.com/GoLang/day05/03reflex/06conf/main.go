package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"reflect"
	"strconv"
	"strings"
)

// 解析日志库文件

// Config  是一个文件配置项
type Config struct {
	FilePath string `conf:"file_Path"`
	FileName string `conf:"file_Name"`
	MaxSize  int    `conf:"max_Size"`
}

func parseConf(fileName string, result interface{}) (err error) {
	// 0. 前提条件  result 必须是个结构体 和指针

	t := reflect.TypeOf(result)
	v := reflect.ValueOf(result)

	if t.Kind() != reflect.Ptr {
		err = errors.New("结构体必须是个指针")
		return
	}

	if t.Elem().Kind() != reflect.Struct {
		err = errors.New("结构体必须是个结构体指针")
		return
	}

	// 1. 打开文件
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		err = errors.New("打开文件失败")
		return err
	}
	// 2. 将读取的文件数据按照行进行分割
	lineSlice := strings.Split(string(data), "\r\n")

	for index, lineStr := range lineSlice {

		lineStr = strings.TrimSpace(lineStr)

		//  如果是空行或者注释 则不遍历数据
		if len(lineStr) == 0 || strings.HasPrefix(lineStr, "#") || strings.HasPrefix(lineStr, ";") {
			continue
		}
		//  解析 判断是不是又等号
		equalIndex := strings.Index(lineStr, "=")
		if equalIndex == -1 {
			err = fmt.Errorf("第%d行有语法错误", index+1)
			return
		}
		//  按照 =  切割

		key := lineStr[:equalIndex]
		value := lineStr[equalIndex+1:]

		key = strings.TrimSpace(key)
		value = strings.TrimSpace(value)

		if len(key) == 0 {
			err = fmt.Errorf("第%d行有语法错误", index+1)
			return
		}

		//  利用反射 给 结构体赋值
		for i := 0; i < t.Elem().NumField(); i++ {
			filed := t.Elem().Field(i)
			tag := filed.Tag.Get("conf")
			if key == tag {
				// 匹配上了 就赋值
				// 拿到每一个字段的类型
				fileType := filed.Type
				// fmt.Println(fileType)
				switch fileType.Kind() {
				case reflect.String:
					// 根据字段名 找到字段值
					// filed.Name 是结构体的Key

					filedValue := v.Elem().FieldByName(filed.Name)
					filedValue.SetString(value)

				case reflect.Int64, reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32:
					value64, _ := strconv.ParseInt(value, 10, 64)
					v.Elem().Field(i).SetInt(value64)

				}

			}

		}

	}

	return
}

func main() {
	// 1. 打开文件

	// 2. 读取内容

	// 3. 一行一行读取内容 根据tag找结构体里面的对应字段

	// 4. 找到了要赋值
	var c = &Config{}
	err := parseConf("xxx.conf", c)
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Println(c)
	fmt.Printf("%#v", c)

}
