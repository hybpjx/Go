package mylog

import (
	"path"
	"runtime"
)

// 存放一些公用的工具函数
func getCallerInfo(skip int) (fileName string, line int, funcName string) {
	// skip 代表跳过多少曾
	pc, fullFileName, line, ok := runtime.Caller(skip)

	if !ok {
		return
	}
	// 从 fullFileName剥离出文件名
	fileName = path.Base(fullFileName)
	// 根据pc拿到函数名 不加name 返回的就是一个func的指针
	funcName = runtime.FuncForPC(pc).Name()
	funcName = path.Base(funcName)
	return

}
