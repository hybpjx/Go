package main

import "code.zic.com/GoLang/day06/homework/mylog"

// 一个使用自定义日志库的程序v
func main() {
	logger := mylog.NewFileLogger("Info", "./", "xx.log")
	defer logger.Close()
	for {
		logger.Debug("Debug 测试")
		logger.Info("Info 测试")
		logger.Error("Error 测试")
		logger.Fatal("Fatal 测试")
	}

}
